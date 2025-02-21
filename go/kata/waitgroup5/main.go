package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func() {

			random := rand.IntN(700) + 200
			time.Sleep(time.Duration(random) * time.Millisecond)

			fmt.Println("id:", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
