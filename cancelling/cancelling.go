package main

import (
	"context"
	"fmt"
	"time"
)

// it is very ordinary form using context which we use regularly
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go PerformTask(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}

func PerformTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			fmt.Println("Performing task...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
