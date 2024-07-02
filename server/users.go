package server

import (
	"net/http"
	"qbc/backend/models"
)

func (s *server) HandleUserList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data []models.User
		s.db.Find(&data)
		s.respond(w, r, data, http.StatusOK)
	}
}

func (s *server) HandleUserRegister() http.HandlerFunc {
	type request struct {
		Email string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var data request
		s.decode(w, r, &data)

		if data.Email == "" {
			http.Error(w, "Email cannot be empty", http.StatusBadRequest)
			return
		}

		newUser := models.User{
			Email: data.Email,
		}
		result := s.db.Create(&newUser)
		if result.Error != nil {
			http.Error(w, "problem occurred", http.StatusInternalServerError)
			return
		}

		signedInUser := models.User{}
		s.db.Where("email = ?", newUser.Email).First(&signedInUser)
		s.respond(w, r, signedInUser, http.StatusCreated)
	}
}
