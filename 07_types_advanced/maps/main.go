package main

import "fmt"

func main() {
	// dictionary
	// hash
	// k<string> => v<int> store
	m := map[string]int{
		"Marcos": 1,
		"Andres": 2,
	}

	dm := make(map[string]string, 100)
	fmt.Println("Len (1):", len(dm))
	// cap(dm)
	fmt.Println(dm)

	fmt.Println("Map [t=0]:", m)

	m["Test"] = 123

	fmt.Println("Map [t=1]:", m)
}
