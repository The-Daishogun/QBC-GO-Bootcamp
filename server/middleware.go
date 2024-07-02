package server

import (
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
