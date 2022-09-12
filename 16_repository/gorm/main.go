package main

import (
	"fmt"
)

type Service struct {
	repository Repository
}

func main() {
	var r Repository

	r = NewInMemoryRepository(nil)

	found, err := r.GetByUsername("marcoshuck2")
	fmt.Println("User:", found)
	fmt.Println("Error:", err)
}
