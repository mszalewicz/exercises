package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	ticker := time.NewTicker(330 * time.Millisecond)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case t := <-ticker.C:
			fmt.Println("tick:", t)
		}
	}

	ticker.Stop()
	fmt.Println("Ticker ended.")
}
