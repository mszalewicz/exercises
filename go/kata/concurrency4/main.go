package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	for i := range 1000 {
		waitGroup.Add(1)


		go func() {
			defer waitGroup.Done()

			randTime := time.Duration(rand.IntN(4050) + 100) * time.Millisecond
			time.Sleep(randTime)

			fmt.Println("goroutine: ", i+1)
		}()

	}

	waitGroup.Wait()
}