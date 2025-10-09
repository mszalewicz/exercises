package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	ticker := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("end")
			return
		}
	}
}
