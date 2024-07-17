package server

import (
	"net/http"
	"qbc/backend/deps"

	"gorm.io/gorm"
)

type server struct {
	db          *gorm.DB
	router      *http.ServeMux
	emailServer deps.EmailSender
}

func NewServer(db *gorm.DB, emailServer deps.EmailSender) *server {
	s := &server{db: db, emailServer: emailServer, router: http.NewServeMux()}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {

	s.router.Handle("GET /users", s.middlewareTimer(s.HandleUserList()))
	s.router.Handle("POST /register", s.middlewareTimer(s.HandleUserRegister()))
	s.router.Handle("POST /send_email", s.middlewareTimer(s.HandleSendEmail()))
	s.router.Handle("GET /calculate/fib/{num}", s.middlewareTimer(s.HandleFib()))

}
