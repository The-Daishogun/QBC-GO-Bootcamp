package server

import (
	"net/http"
	"qbc/backend/deps"

	"gorm.io/gorm"
)

type server struct {
	db          *gorm.DB
	router      *http.ServeMux
	emailServer *deps.EmailServer
	caches      *deps.Caches
}

func NewServer(db *gorm.DB, emailServer *deps.EmailServer) *server {
	s := &server{db: db, emailServer: emailServer, router: http.NewServeMux(), caches: &deps.Caches{
		Calculations: deps.NewRedisClient(0),
		Responses:    deps.NewRedisClient(1),
	}}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {

	s.router.Handle("GET /users", s.middlewareTimer(s.HandleUserList()))
	s.router.Handle("POST /user", s.middlewareTimer(s.HandleUserRegister()))
	s.router.Handle("POST /send_email", s.middlewareTimer(s.HandleSendEmail()))
	s.router.Handle("GET /calculate/fib/{num}", s.middlewareTimer(s.middlewareCacheResponseRequestURI(s.HandleFib())))

}
