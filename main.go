package main

import (
	"log"
	"net/http"
	"qbc/backend/deps"
	"qbc/backend/server"
	"time"
)

func main() {
	MaxConcurrentWorkers := 5

	log.Println("Setting Up the Server...")
	db := deps.CreateNewDB()
	emailServer := deps.NewEmailServer(MaxConcurrentWorkers, time.NewTicker(500*time.Millisecond))
	go emailServer.RunPostOffice()
	s := server.NewServer(db, emailServer)
	log.Println("HTTP Server Running on http://localhost:8000")
	err := http.ListenAndServe(":8000", s)
	if err != nil {
		log.Fatal("Error Starting server. Error: ", err.Error())
	}
}
