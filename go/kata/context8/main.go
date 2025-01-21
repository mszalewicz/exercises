package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChannel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	defer cancel()

	go ThirdPartyQuery(ctx, responseChannel)

	select {
		case <-ctx.Done():
			fmt.Println("third party query timed out")
		case respone := <-responseChannel:
			fmt.Println(respone)
	}
}

func ThirdPartyQuery(ctx context.Context, responseChannel chan string) {
	// ping

	for {
		select {
			case <-ctx.Done():
				return
			default:
				// check if third party responded
				time.Sleep(100 * time.Millisecond)
				responseChannel <- "third party info: {...}"
		}
	}
}