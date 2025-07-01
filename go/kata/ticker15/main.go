package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(350 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

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
