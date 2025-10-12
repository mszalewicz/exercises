package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	counter := 0

	for i := range 2000 {
		go func(counter *int) {
			wg.Add(1)

			rnd := time.Duration(rand.Int32N(8400)) * time.Millisecond + 500 * time.Millisecond
			time.Sleep(rnd)
			fmt.Println("goroutine:", i)

			mutex.Lock()
			*counter++
			mutex.Unlock()


			wg.Done()
		}(&counter)
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
