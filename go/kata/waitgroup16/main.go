package main

import (
	"math/rand/v2"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			rndTime := time.Duration(rand.Int32N(400)) * time.Millisecond + 140 * time.Millisecond
			time.Sleep(rndTime)
			fmt.Println("goroutine:", i)
		}(i + 1)
	}

	wg.Wait()
}