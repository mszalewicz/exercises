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

	for i := range 1000 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			rndTime := time.Duration(rand.Int32N(800))*time.Millisecond + 100*time.Millisecond
			time.Sleep(rndTime)

			mutex.Lock()
			counter++
			mutex.Unlock()

			fmt.Println(time.Now())
		}(i)
	}

	wg.Wait()
	fmt.Println("Goroutines done:", counter)
}
