package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(400 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("Job done!")
			return
		}
	}
}