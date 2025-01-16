package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChannel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 150 * time.Millisecond)
	defer cancel()

	go ThirdPartyQuery(ctx, responseChannel)

	select {
		case <-ctx.Done():
			fmt.Println("Third party query timed out.")
		case response := <-responseChannel:
			fmt.Println(response)
	}
}

func ThirdPartyQuery(ctx context.Context, responseChannel chan string) {
	for {
		select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(200 * time.Millisecond)
				responseChannel <- "Third party info: {...}"
		}
	}
}