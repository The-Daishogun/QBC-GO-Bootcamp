package main

import (
	"fmt"
	"time"
)

func work(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	work("direct")

	go work("goroutine")

	go func() {
		fmt.Println("going")
	}()

	time.Sleep(time.Second)
	fmt.Println("done")
}
