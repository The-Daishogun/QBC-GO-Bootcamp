package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *server) middlewareTimer(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received\t%s\t%s\n", req.Method, req.URL.Path)
		startTime := time.Now()

		h(w, req)

		duration := time.Since(startTime)
		log.Printf("Finished\t%s\t%s\t%s\n", req.Method, req.URL.Path, duration)
	}
}

type cachedResponse struct {
	Body   string
	Status int
}

type responseRecorder struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
		body:           &bytes.Buffer{},
		statusCode:     http.StatusOK,
	}
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (s *server) middlewareCacheResponseRequestURI(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		cacheKey := r.RequestURI

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cachedResponseStr, err := s.caches.Responses.Get(ctx, cacheKey).Result()
		if err == nil {
			w.Header().Set("Response-Cache-Status", "HIT")
			var cachedResponse cachedResponse
			json.Unmarshal([]byte(cachedResponseStr), &cachedResponse)
			s.respond(w, r, cachedResponse.Body, cachedResponse.Status)
			return
		}
		log.Println("CACHE MISS\t", fmt.Sprintf("%s\t", r.RequestURI), cacheKey)

		recorder := newResponseRecorder(w)
		h(recorder, r)

		responseBody := recorder.body.String()
		responseStatus := recorder.statusCode
		cacheData := cachedResponse{
			Body:   responseBody,
			Status: responseStatus,
		}
		cacheDataBytes, _ := json.Marshal(cacheData)
		_, err = s.caches.Responses.Set(ctx, cacheKey, cacheDataBytes, 120*time.Second).Result()
		if err != nil {
			log.Println("failed to set cache for handler ", err)
		}
		w.Header().Set("Response-Cache-Status", "MISS")
	}
}
