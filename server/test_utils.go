package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qbc/backend/deps"

	"github.com/stretchr/testify/mock"
)

type mockedEmailServer struct {
	mock.Mock
}

func (m *mockedEmailServer) SendEmail(to, subject, content string) {
	m.Called(to, subject, content)
}

func newTestServer() (*server, func()) {
	db, tearDownDB := deps.CreateNewDB("test.db")

	emailServer := new(mockedEmailServer)
	s := NewServer(db, emailServer)

	return s, func() {
		tearDownDB()
	}
}

func setupTestSuiteWithRequest(tc any, method, endpoint string) (*server, *httptest.ResponseRecorder, func()) {
	s, tearDown := newTestServer()

	bodyBytes, _ := json.Marshal(tc)
	bodyReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest(method, endpoint, bodyReader)

	rr := httptest.NewRecorder()

	s.router.ServeHTTP(rr, req)

	return s, rr, func() {
		tearDown()
	}
}
