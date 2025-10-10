package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ticker := time.NewTicker(200 * time.Millisecond)

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
