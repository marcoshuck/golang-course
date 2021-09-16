package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		fmt.Println("Chat two ---")
		go func() {
			for {
				one <- "Testing message"
				time.Sleep(5 * time.Second)
			}
		}()
		for m := range two {
			log.Println("One:", m)
		}
	}()

	fmt.Println("Chat one ---")

	for m := range one {
		log.Println("Two:", m)
	}

	for {
		var msg string
		_, err := fmt.Scanf("%s", &msg)
		if err != nil {
			log.Println("Error sending the message")
			continue
		}

		two <- msg
	}
}
