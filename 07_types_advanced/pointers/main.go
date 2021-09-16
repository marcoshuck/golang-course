package main

import "fmt"

func main() {
	a := 42

	var inputA int
	fmt.Println("Ingresar primer valor:")
	fmt.Scan(&inputA)

	p := &a // C0000012112125 (*int)

	fmt.Println("Value:", a)
	fmt.Println("Pointer:", p)

	*p = 24

	fmt.Println("Value:", a)
	fmt.Println("Pointer:", p)

	A(a)
	fmt.Println("Call A:", a)

	B(&a)
	fmt.Println("Call B:", a)
}

func A(a int) {
	a = 25
}

func B(b *int) {
	*b = 25
}
