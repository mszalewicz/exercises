package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 4000 {
		wg.Add(1)

		go func() {
			randTime := time.Duration(rand.IntN(850) * int(time.Millisecond))
			time.Sleep(randTime)

			fmt.Println("id: ", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
