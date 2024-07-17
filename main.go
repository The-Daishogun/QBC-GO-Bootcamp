package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"qbc/backend/deps"
	"qbc/backend/server"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "The port to run the server on")
	flag.Parse()
}

func main() {
	log.Println("Setting Up the Server...")
	db, _ := deps.CreateNewDB("db.sqlite")
	emailServer := deps.NewEmailServer()
	s := server.NewServer(db, emailServer)
	log.Printf("HTTP Server Running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s)
	if err != nil {
		log.Fatal("Error Starting server. Error: ", err.Error())
	}
}
