package main

import (
	"log"
	"time"
)

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	log.Println("sent buffered")
	messages <- "channel"
	log.Println("sent channel")

	log.Println(<-messages)
	time.Sleep(1 * time.Second)
	log.Println(<-messages)
}
