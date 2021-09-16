package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key string

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, ok := ctx.Value("request_id").(string)
		if ok {
			log.Println("Request ID:", id)
		}

		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()

		select {
		// Process context
		case <-ctx.Done():
			// Rollback.
			log.Println("Context error:", ctx.Err())
			http.Error(w, fmt.Sprintf("Cancelled: %s", ctx.Err()), http.StatusInternalServerError)

		// Normal execution
		case <-time.After(10 * time.Second):
			w.WriteHeader(http.StatusOK)
			_, err := fmt.Fprintln(w, "Hello, world!")
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
			}
			fmt.Println("Hello, world from server!")
			// Requesting a machine
		}
	})

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println("Shutdown")
}
