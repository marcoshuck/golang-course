package main

import "fmt"

func main() {
	messages := make(chan int)

	for i := 0; i < 100; i++ {
		go func(index int) {
			messages <- index
		}(i)
	}

	for i := 0; i < 101; i++ {
		msg := <-messages
		fmt.Println(msg)
	}

}
