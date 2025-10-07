package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	response := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	go query(ctx, response)

	select {
	case <-ctx.Done():
		fmt.Println("timed out")
	case msg := <-response:
		fmt.Println(msg)
	}
}

func query(ctx context.Context, response chan<- string) {
	time.Sleep(50 * time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "info"
	}
}