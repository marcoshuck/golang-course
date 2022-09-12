package main

import "fmt"

func main() {
	// 0.0000000000000000000001
	// 0.0000000000000000000000000000000000000011
	var value32 float32 // float
	var value64 float64 // double

	value32 = 1.0
	value64 = 2.0

	var result float64
	result = float64(value32) + value64

	// ${value32} + ${value64} = ${result}
	fmt.Printf("%.2f + %.2f = %.2f", value32, value64, result)
	// 1.00 + 2.00 = 3.00
}
