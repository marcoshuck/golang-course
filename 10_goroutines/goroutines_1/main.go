package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(1 * time.Second)
}

func test() {
	fmt.Println("Test!")
}
