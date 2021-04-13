package main

import (
	"log"
	"time"
)

func main() {
	// max
	messages := make(chan int)

	go func() {
		log.Println("Attempting to add an element in 1")
		messages <- 1
		log.Println("Added an element in 1")
	}()

	go func() {
		log.Println("Attempting to add an element in 2")
		messages <- 2
		log.Println("Added an element in 2")
	}()

	time.Sleep(5 * time.Second)
	log.Println("Reading a message from the channel:", <-messages)

	time.Sleep(5 * time.Second)
	log.Println("Reading a message from the channel:", <-messages)
}
