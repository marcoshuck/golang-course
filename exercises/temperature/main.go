package main

import "fmt"

const k = 5 / 9

func convertToCelsius(f float64) float64 {
	return ((f - 32) * 5) / 9
}

func main() {

	var fahrenheit float64
	fmt.Println("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scanf("%f\n", &fahrenheit)

	celsius := convertToCelsius(fahrenheit)

	fmt.Printf("%f Fahrenheit equivalen a %f° Celsius.\n", fahrenheit, celsius)

	//var fahrenheit float64
	//var celsius float64
	//fmt.Println("Ingrese la temperatura en Fahrenheit: ")
	//fmt.Scanf("%f\n", &fahrenheit)
	//celsius = (5 / 9) * (fahrenheit - 32)
	//fmt.Printf("%f Fahrenheit equivalen a %v° Celsius.\n", fahrenheit, celsius)
}
