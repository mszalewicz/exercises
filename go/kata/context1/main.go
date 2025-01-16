package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChan := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 200)
	defer cancel()

	go getThirdPartyInfo(responseChan, ctx)

	select {
		case <- ctx.Done():
			fmt.Println("Third party query timed out.")
		case response := <- responseChan:
			fmt.Println(response)
	}
}

func getThirdPartyInfo(response chan string, ctx context.Context) {
	for {
		select {
			case <- ctx.Done():
				return
			default:
				// pinging the third party api
				 time.Sleep(time.Millisecond * 1000)
				response <- "This is third party info."
		}
	}
}