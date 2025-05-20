package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChannel := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Millisecond)
	defer cancel()

	go ThirdPartyInfo(ctx, responseChannel)

	select {
	case <-ctx.Done():
		fmt.Println("3rd party query timed out")
	case response := <-responseChannel:
		fmt.Println(response)
	}

}

func ThirdPartyInfo(ctx context.Context, responseChannel chan string) {
	time.Sleep(50 * time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		responseChannel <- "3rd party info: {...}"
	}
}
