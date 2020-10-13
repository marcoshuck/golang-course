package main

import "fmt"

func main() {
	a := 42

	p := &a

	fmt.Println("Value:", a)
	fmt.Println("Pointer:", p)

	*p = 24

	fmt.Println("Value:", a)
	fmt.Println("Pointer:", p)
}
