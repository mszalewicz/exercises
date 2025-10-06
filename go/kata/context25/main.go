package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChannel := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go thirdPartyQuery(ctx, responseChannel)

	select {
	case <-ctx.Done():
		fmt.Println("query timed out")
	case response := <-responseChannel:
		fmt.Println(response)
	}
}

func thirdPartyQuery(ctx context.Context, responseChannel chan string) {
	time.Sleep(50 * time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		responseChannel <- "3rd party info"
	}
}
