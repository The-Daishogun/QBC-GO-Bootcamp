package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qbc/backend/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	tests := []struct {
		to, subject, content string
	}{
		{"u1@quera.org", "hello", "how are you?"},
	}

	for _, tc := range tests {
		t.Run("testSendingEmail", func(t *testing.T) {
			s, tearDown := newTestServer()
			defer tearDown()
			mockedEmailServer, _ := s.emailServer.(*mockedEmailServer)
			mockedEmailServer.On("SendEmail", tc.to, tc.subject, tc.content)
			s.emailServer.SendEmail(tc.to, tc.subject, tc.content)
			mockedEmailServer.AssertExpectations(t)
		})
	}
}

func TestSendEmailEndpoint(t *testing.T) {
	test := struct {
		Subject string
		Content string
	}{
		Subject: "hello",
		Content: "how are you?",
	}

	s, tearDown := newTestServer()
	defer tearDown()

	mockedEmailServer, _ := s.emailServer.(*mockedEmailServer)
	emails := []string{
		"u1@quera.org",
		"u2@quera.org",
		"u3@quera.org",
	}

	for _, email := range emails {
		s.db.Create(&models.User{
			Email: email,
		})
		mockedEmailServer.On("SendEmail", email, test.Subject, test.Content)
	}

	bodyBytes, _ := json.Marshal(test)
	bodyReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest(http.MethodPost, "/send_email", bodyReader)

	rr := httptest.NewRecorder()

	s.router.ServeHTTP(rr, req)

	mockedEmailServer.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rr.Code)
}
