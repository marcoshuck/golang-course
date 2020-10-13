package main

import "fmt"

func main() {
	m := map[string]int{
		"Marcos":  1,
		"Gabriel": 2,
	}

	fmt.Println("Map [t=0]:", m)

	m["Test"] = 123

	fmt.Println("Map [t=1]:", m)
}
