package main

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

const userIDKey contextKey = "UserID"

// in this code I sent context to goroutines with value
func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, userIDKey, 123)

	go PerformTask(ctx)

	time.Sleep(1 * time.Second)
}

func PerformTask(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	fmt.Println("User ID:", userID)
}
