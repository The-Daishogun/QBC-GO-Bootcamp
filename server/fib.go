package server

import (
	"net/http"
	"strconv"
)

func calculate_fib(num int) int {
	if num <= 1 {
		return num
	}

	return calculate_fib(num-2) + calculate_fib(num-1)
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
		answer := calculate_fib(num)

		w.Header().Set("Cache-Control", "max-age=31536000") // Cache the result for a year
		s.respond(w, r, response{Answer: answer}, http.StatusOK)
	}
}
