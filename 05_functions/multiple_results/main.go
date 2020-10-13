package main

import "fmt"

func main() {
	a, b := swap("world", "hello")
	fmt.Println(a, b)
}

func swap(a, b string) (string, string) {
	return b, a
}
