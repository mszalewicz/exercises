package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++{
		go func() {
			wg.Add(1)

			rndTime := time.Duration(rand.IntN(500) + 150) * time.Millisecond
			time.Sleep(rndTime)

			fmt.Println("Goroutine ", i)

			wg.Done()
		}()
	}

	wg.Wait()
}