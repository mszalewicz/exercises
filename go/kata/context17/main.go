package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	responseChannel := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	go func(responseChannel chan string) {
		time.Sleep(20 * time.Millisecond)

		responseChannel <- "3rd party example info: ..."
	}(responseChannel)

	select {
	case <-ctx.Done():
		fmt.Println("3rd party query timed out")
	case response := <-responseChannel:
		fmt.Println(response)
	}
}
