package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var mutext sync.Mutex
	var wg sync.WaitGroup

	counter := 0

	for i := range 1000 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			rndTime := time.Duration(rand.Int32N(3500))*time.Millisecond + 100*time.Millisecond
			time.Sleep(rndTime)

			mutext.Lock()
			fmt.Println("goroutine:", i)
			counter++
			mutext.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("all goroutines = ", counter)
}
