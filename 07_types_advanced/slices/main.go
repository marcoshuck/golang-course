package main

import "fmt"

func main() {
	// var arr [3]int !=
	var slice []int

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	slice = append(slice, 100)
	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	anotherSlice := make([]int, 0, 500)
	fmt.Println("Size:", len(anotherSlice))
	fmt.Println("Capacity:", cap(anotherSlice))

	var arr [5]int
	fmt.Println("Size:", len(arr))
	fmt.Println("Capacity:", cap(arr))

	length := 10
	v := make([]int, length)
	v[0] = 1

	fmt.Println("Size:", len(v))
	fmt.Println("Capacity:", cap(v))

	s := make([]int, 10, 10)
	fmt.Println(s)

	s = append(s, 120)
	fmt.Println("Size:", len(s))
	fmt.Println("Capacity:", cap(s))
}
