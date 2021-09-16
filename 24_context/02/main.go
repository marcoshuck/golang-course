package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Execute a process
	err := process(ctx, 1)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// Finish execution
	log.Println("Success")
}

func process(ctx context.Context, value int) error {
	select {
	case <-ctx.Done():
		log.Println("Cancelled")
		return ctx.Err()
	case <-time.After(5 * time.Second):
		log.Println(fmt.Sprintf("Value %d + 1 equals %d", value, value+1))
		log.Println("Normal execution")
	}
	return nil
}
