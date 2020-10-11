package main

import "fmt"

func main() {
	msg := "Hello, World"

	fmt.Println(msg[0])

	fmt.Println(len(msg))

	fmt.Println(fmt.Sprintf("%s. We're ready!", msg))
}
