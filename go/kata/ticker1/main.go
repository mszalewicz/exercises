package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(400 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				fmt.Println("Ticker:", t)
			}
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Ticker stopped.")
		ticker.Stop()
		return
	}
}
