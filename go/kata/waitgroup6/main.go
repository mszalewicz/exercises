package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for index := range 15 {
		wg.Add(1)

		go func() {
			rnd := rand.IntN(500) + 200
			time.Sleep(time.Duration(rnd) * time.Millisecond)
			fmt.Println("index:", index+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
