package server

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string
}

func (s *server) decode(_ http.ResponseWriter, r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data == nil {
		return
	}

	strData, ok := data.(string)
	if ok {
		w.Write([]byte(strData))
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		s.respond(w, r, nil, http.StatusInternalServerError)
	}
}
