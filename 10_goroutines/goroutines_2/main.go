package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int
	go func() {
		sum++
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(sum)
}
