package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	responseChannel := make(chan string)

	go thirdPartyQuery(ctx, responseChannel)

	select {
	case <-ctx.Done():
		fmt.Println("Third party query timed out.")
	case response := <-responseChannel:
		fmt.Println(response)
	}

}

func thirdPartyQuery(ctx context.Context, responseChannel chan string) {
	time.Sleep(400 * time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		responseChannel <- "third party info: ()"
	}
}
