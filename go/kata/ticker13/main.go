package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ticker := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("Ticker stopped.")
			os.Exit(0)
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}
