package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func CalculateFibWithCaching(num uint64, cache *redis.Client) uint64 {
	if num <= 1 {
		return num
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cacheKey := fmt.Sprintf("fib_%d", num)
	fibStr, err := cache.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("CACHE HIT")
		fibInt, _ := strconv.ParseUint(fibStr, 10, 64)
		return fibInt
	}

	fib := CalculateFibWithCaching(num-2, cache) + CalculateFibWithCaching(num-1, cache)

	_, err = cache.Set(ctx, cacheKey, fib, 0).Result()
	if err != nil {
		log.Printf("failed to set cache for %s, Error: %s\n", cacheKey, err.Error())
	}
	return fib
}

func CaclulateFib(num uint64) uint64 {
	if num <= 1 {
		return num
	}
	return CaclulateFib(num-2) + CaclulateFib(num-1)
}

func (s *server) HandleFib() http.HandlerFunc {
	type response struct {
		Answer uint64
	}
	return func(w http.ResponseWriter, r *http.Request) {
		num, err := strconv.ParseUint(r.PathValue("num"), 10, 64)
		if err != nil || num == 0 {
			s.respond(w, r, ErrorResponse{Error: "Invalid Number"}, http.StatusBadRequest)
			return
		}
		answer := CalculateFibWithCaching(num, s.caches.Calculations)

		w.Header().Set("Cache-Control", "max-age=31536000") // Cache the result for a year
		s.respond(w, r, response{Answer: answer}, http.StatusOK)
	}
}
