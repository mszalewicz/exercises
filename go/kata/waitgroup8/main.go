package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 120 {
		wg.Add(1)

		go func() {
			rndSleepTime := rand.IntN(1750) + 200
			time.Sleep(time.Duration(rndSleepTime) * time.Millisecond)

			fmt.Println("id: ", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
