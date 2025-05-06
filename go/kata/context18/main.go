package main

import (
	"context"
	"fmt"
	"time"
)

func getThirdParty(ctx context.Context, responseChannel chan string) {
	time.Sleep(110 * time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		responseChannel <-"3rd party info: ..."
		return
	}
}

func main() {
	responseChannel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	go getThirdParty(ctx, responseChannel)

	select {
	case response := <-responseChannel:
		fmt.Println(response)
	case <-ctx.Done():
		fmt.Println("3rd party query timed out")
	}
}