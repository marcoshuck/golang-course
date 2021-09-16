package main

import (
	"context"
	"fmt"
)

type key string

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, key("test"), "1234")

	printContextValue(ctx, key("test"))
}

func printContextValue(ctx context.Context, k key) {
	fmt.Println(ctx.Value(k))
}
