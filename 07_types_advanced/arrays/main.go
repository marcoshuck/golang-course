package main

import "fmt"

func main() {
	var array [3]int

	fmt.Println("Size:", len(array))

	fmt.Println("Array [t=0]:", array)

	for i := 0; i < len(array); i++ {
		array[i] = 3 * i
	}

	fmt.Println("Array [t=1]:", array)
}
