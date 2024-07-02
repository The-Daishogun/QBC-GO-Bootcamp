package main

import (
	"log"
	"net/http"
	"qbc/backend/deps"
	"qbc/backend/server"
)

func main() {
	log.Println("Setting Up the Server...")
	db := deps.CreateNewDB()
	emailServer := deps.NewEmailServer()
	s := server.NewServer(db, emailServer)
	log.Println("HTTP Server Running on http://localhost:8000")
	err := http.ListenAndServe(":8000", s)
	if err != nil {
		log.Fatal("Error Starting server. Error: ", err.Error())
	}
}
