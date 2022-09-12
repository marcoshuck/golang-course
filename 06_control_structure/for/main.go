package main

import "fmt"

func main() {
	// initial setting; condition; step;
	for i := 1; i < 10; i = i + 1 {
		fmt.Printf("%d ", i)
	}
}
