package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	queue = make(chan string, 1_000_000)
)

func main() {
	go func() {
		for ip := range queue {
			fmt.Println("IP Received:", ip)
			time.Sleep(15 * time.Second)
		}
	}()

	if err := http.ListenAndServe(":8080", http.HandlerFunc(insert)); err != nil {
		panic(err)
	}
}

func insert(w http.ResponseWriter, r *http.Request) {
	queue <- r.RemoteAddr
	w.WriteHeader(http.StatusNoContent)
}
