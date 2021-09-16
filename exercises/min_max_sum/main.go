package main

import "fmt"

func main() {
	arr := []int{7, 69, 2, 221, 8974}

	//var ordered []int
	Order(arr)

	//minSum := getMinSum(arr)
	//maxSum := getMaxSum(arr)
}

func Order(arr []int) []int {
	var ordered []int

	// O(n^2)
	// [2, 69, 7, 221, 8974]

	fmt.Println("Ordered:", arr)

	return ordered
}

func getMaxSum(arr []int) int {
	var sum int
	for i := 0; i < len(arr)-1; i++ {
		sum += arr[i]
	}
	return sum
}

func getMinSum(arr []int) int {
	var sum int
	for i := 1; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
