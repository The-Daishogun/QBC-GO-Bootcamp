package main

import (
	"log"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pong"
		log.Println("rast migam?")
	}()

	msg := <-messages
	log.Println(msg)
	time.Sleep(time.Second)
	msg2 := <-messages
	log.Println(msg2)
	time.Sleep(time.Second)

}
