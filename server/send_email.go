package server

import (
	"fmt"
	"net/http"
	"qbc/backend/models"
)

func (s *server) HandleSendEmail() http.HandlerFunc {
	type request struct {
		Subject string
		Content string
	}
	type response struct {
		Message string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		data := request{}
		s.decode(w, r, &data)

		if data.Subject == "" || data.Content == "" {
			s.respond(w, r, ErrorResponse{Error: "Subject, Content should not be empty"}, http.StatusBadRequest)
			return
		}

		var allUsers []models.User
		s.db.Find(&allUsers)

		if allUsers == nil {
			s.respond(w, r, ErrorResponse{Error: "no users found"}, http.StatusBadRequest)
			return
		}
		for _, user := range allUsers {
			s.emailServer.SendEmail(user.Email, data.Subject, data.Content)
		}

		s.respond(w, r, response{Message: fmt.Sprintf("%d emails sent!", len(allUsers))}, http.StatusOK)
	}
}
