package main

import (
	"fmt"
)

func main() {
	sum := 1

	// for [Condicion inicial]; [Comprobacion de condicion]; [Incremento o proceso por iteración] { }

	for sum < 1000 {
		sum += sum
	}

	fmt.Println("Sum:", sum)
}
