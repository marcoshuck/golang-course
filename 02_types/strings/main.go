package main

import "fmt" // format

func main() {
	var msg string
	msg = "Hello, World"

	// ASCII 72 number = H
	fmt.Println(msg[0])

	fmt.Println(len(msg))

	// ${} .{}
	newMsg := fmt.Sprintf("%s. We're ready!", msg)
	fmt.Println(newMsg)
}
