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
			howLong := rand.IntN(300) + 50
			time.Sleep(time.Duration(howLong) * time.Millisecond)

			fmt.Println("goroutine: ", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
