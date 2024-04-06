package main

import (
	"context"
	"fmt"
	"time"
)

// in this code I can set deadline for context, actually it works very simple, I guess
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	done := make(chan struct{})

	go PerformTask(ctx, done)

	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Deadline exceeded:", ctx.Err())
	}

	time.Sleep(1 * time.Second)
}

func PerformTask(ctx context.Context, done chan<- struct{}) {
	time.Sleep(3 * time.Second)

	done <- struct{}{}
}
