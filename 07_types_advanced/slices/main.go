package main

import "fmt"

func main() {
	var slice []int

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	anotherSlice := make([]int, 0, 500)

	length := 1

	var arr [5]int

	v := make([]int, length)

	v[0] = 1

	fmt.Println("Size:", len(anotherSlice))
	fmt.Println("Capacity:", cap(anotherSlice))
}
