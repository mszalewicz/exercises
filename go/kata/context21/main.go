package main

import (
	"context"
	"time"
	"fmt"
)

func main() {
	responseChannel := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
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
	select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(150 * time.Millisecond)
			responseChannel <-"Third party info: {...}"
	}
}