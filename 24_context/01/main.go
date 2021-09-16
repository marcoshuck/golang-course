package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	// Create context
	ctx := context.Background()

	// Execute a process
	err := process(ctx, 1)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// Finish execution
	log.Println("Success")
}

// AWS SDK for Go

func process(ctx context.Context, value int) error {
	log.Println(fmt.Sprintf("Value %d + 1 equals %d", value, value+1))
	return nil
}
