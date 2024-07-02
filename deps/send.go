package deps

import (
	"log"
	"time"
)

type EmailServer struct {
}

func (es *EmailServer) SendEmail(to, subject, content string) {
	time.Sleep(2000 * time.Millisecond)
	log.Println("Email Sent!")
}

func NewEmailServer() *EmailServer {
	return &EmailServer{}
}
