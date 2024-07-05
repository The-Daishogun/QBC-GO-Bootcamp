package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func calculate_fib(num int, cache *redis.Client) int {
	if num <= 1 {
		return num
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cacheKey := fmt.Sprintf("fib_%d", num)
	fibStr, err := cache.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("CACHE HIT")
		fibInt, _ := strconv.Atoi(fibStr)
		return fibInt
	}

	fib := calculate_fib(num-2, cache) + calculate_fib(num-1, cache)

	_, err = cache.Set(ctx, cacheKey, fib, 0).Result()
	if err != nil {
		log.Printf("failed to set cache for %s, Error: %s\n", cacheKey, err.Error())
	}
	return fib
}

func (s *server) HandleFib() http.HandlerFunc {
	type response struct {
		Answer int
	}
	return func(w http.ResponseWriter, r *http.Request) {
		num, err := strconv.Atoi(r.PathValue("num"))
		if err != nil || num < 0 {
			s.respond(w, r, ErrorResponse{Error: "Invalid Number"}, http.StatusBadRequest)
			return
		}
		answer := calculate_fib(num, s.caches.Calculations)

		w.Header().Set("Cache-Control", "max-age=31536000") // Cache the result for a year
		s.respond(w, r, response{Answer: answer}, http.StatusOK)
	}
}
