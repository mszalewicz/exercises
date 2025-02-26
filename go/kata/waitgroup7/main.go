package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 50 {
		wg.Add(1)

		go func() {
			randomTime := rand.IntN(6000) + 250
			time.Sleep(time.Duration(randomTime) * time.Millisecond)

			fmt.Println("id: ", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
