package main

import (
	"fmt"
	"log"
)

func main() {

	var inputA int
	fmt.Println("Ingresar primer valor:")
	_, err := fmt.Scan(&inputA)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Input A:", inputA)

	value := 42
	// getting the reference (hex code) from memory
	p := &value // C0000012112125 (*int)

	fmt.Println("Value:", value)
	fmt.Println("Pointer:", p)

	// dereference the address
	// value = 24
	*p = 24

	fmt.Println("Value:", value)
	fmt.Println("Pointer:", p)
	fmt.Println("Pointer of pointer:", &p)

	A(value) // value = 24
	fmt.Println("Call A:", value)

	b := &value
	B(b)
	fmt.Println("Call B:", value)
}

func A(a int) {
	a = 25
}

func B(b *int) {
	*b = 25
}
