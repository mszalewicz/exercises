package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	responseChannel := make(chan string)

	go thirdPartyQuery(ctx, responseChannel)

	select {
	case <-ctx.Done():
		fmt.Println("query timed out")
	case response := <-responseChannel:
		fmt.Println("thir party response: ", response)
	}
}

func thirdPartyQuery(ctx context.Context, responseChannel chan string) {
	time.Sleep(5 * time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		responseChannel <- "+ 3rd party info +"
	}
}
