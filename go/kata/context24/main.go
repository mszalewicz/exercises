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

	go thirdPartyInfo(ctx, responseChannel)

	select {
	case <-ctx.Done():
		fmt.Println("3rd party query timed out")
		return
	case response := <-responseChannel:
		fmt.Println(response)
	}
}

func thirdPartyInfo(ctx context.Context, responseChannel chan string) {
	// simulate query time execution
	time.Sleep(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			responseChannel <- "3rd party info {...}"
		}
	}
}
