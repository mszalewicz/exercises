package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChannel := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	go ThirdPartyQuery(ctx, responseChannel)

	select {
		case <- ctx.Done():
			fmt.Println("Third party query timed out.")
		case response := <-responseChannel:
			fmt.Println(response)
	}
}

func ThirdPartyQuery(ctx context.Context, responseChannel chan string) {

	// ping 3rd party

	for {
		select {
			case <-ctx.Done():
				return
			default:
				// check if there was a response from 3rd party
				time.Sleep(150 * time.Millisecond)
				responseChannel <- "Third party data ..."
		}
	}
}
