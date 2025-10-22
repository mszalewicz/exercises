package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	response := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
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
	rndTime := time.Duration(rand.Int32N(100))*time.Millisecond + 250*time.Millisecond
	time.Sleep(rndTime)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "query result"
	}
}
