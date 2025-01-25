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
			time.Sleep(time.Duration(rand.IntN(1500)+300) * time.Millisecond)
			fmt.Println("goroutine:", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
}
