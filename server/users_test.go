package server

import (
	"fmt"
	"net/http"
	"qbc/backend/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		Email      string `json:"email"`
		statusCode int
	}{
		{"javad@quera.org", http.StatusCreated},
		{"", http.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("registering %s", tc.Email), func(t *testing.T) {
			s, rr, tearDownFunc := setupTestSuiteWithRequest(tc, http.MethodPost, "/register")
			defer tearDownFunc()

			assert.Equal(t, tc.statusCode, rr.Code, fmt.Sprintf("handler return wrong status code for %s. wanted: %d, returned: %d", tc.Email, tc.statusCode, rr.Code))

			if tc.statusCode == 201 {

				var exists bool
				s.db.Model(&models.User{}).
					Select("count(*) > 0").
					Where("email = ?", tc.Email).
					Find(&exists)
				assert.True(t, exists, fmt.Sprintf("user with email %s does not exists in the db", tc.Email))
			}
		})
	}
}
