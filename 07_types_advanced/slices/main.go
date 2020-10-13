package main

import "fmt"

func main() {
	var slice []int

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	slice = append(slice, 15)

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))
}
