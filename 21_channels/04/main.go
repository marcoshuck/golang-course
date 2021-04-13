package main

import (
	"fmt"
	"time"
)

func main() {
	ready := make(chan bool)

	go func() {
		fmt.Println("Hey, I'm doing heavy stuff here, wait for me.")
		time.Sleep(time.Second)
		ready <- true
	}()

	<-ready
}
