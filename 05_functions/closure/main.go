package main

import "fmt"

func generateCounter() (func(), *int) {
	var count int
	return func() { count++ }, &count
}

func main() {
	counter, count := generateCounter()

	for i := 0; i < 100; i++ {
		counter()
	}

	fmt.Println("Count:", *count)
}
