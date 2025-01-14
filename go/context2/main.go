package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChan := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 200 * time.Millisecond)
	defer cancel()

	go ThirdPartyQuery(ctx, responseChan)

	select {
		case response := <- responseChan:
			fmt.Println(response)
		case <- ctx.Done():
			fmt.Println("Third party query timed out.")
	}
}

func ThirdPartyQuery(ctx context.Context, responseChan chan string) {
	time.Sleep(100 * time.Millisecond)

	for {
		select {
			case <- ctx.Done():
				return
			default:
				time.Sleep(300 * time.Millisecond)
				responseChan <- "Third party information..."
		}
	}
}