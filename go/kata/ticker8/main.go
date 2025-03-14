package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(30 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("ticker stopped")
			return
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}