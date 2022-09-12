package main

import "fmt"

func main() {
	var slice []int

	for i := 0; i < 100; i++ {
		slice = append(slice, i*100)
	}

	for i := 0; i < len(slice); i++ {

	}

	// foreach loop
	// index, value := range collection
	for i, value := range slice {
		fmt.Println(i, ":", value)
	}
}
