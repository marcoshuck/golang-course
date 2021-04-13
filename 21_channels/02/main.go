package main

import (
	"fmt"
)

func main() { // 1
	messages := make(chan string)

	messages <- "Test"

	msg := <-messages

	fmt.Println(msg)
}
