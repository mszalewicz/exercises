package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(300 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("ticker stopped")
			return
		}
	}
}