package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(value int) {
			fmt.Println(value)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
