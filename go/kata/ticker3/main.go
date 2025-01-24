package main	

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	ticker := time.NewTicker(135 * time.Millisecond)

	go func() {
		for {
			select {
			case tick := <-ticker.C:
				fmt.Println(tick)	
			case <-ctx.Done():
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		ticker.Stop()
		fmt.Println("Ticker stopped.")
	}
}
