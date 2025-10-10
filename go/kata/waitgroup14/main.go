package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutext sync.Mutex
	counter := 0

	for i := range 100 {
		wg.Add(1)
		go func(counter *int, mutext *sync.Mutex) {
			rndTime := time.Duration(rand.Int32N(800) + 300) * time.Millisecond
			time.Sleep(rndTime)

			mutext.Lock()
			*counter++
			mutext.Unlock()

			fmt.Println("thread:", i)
			wg.Done()

		}(&counter, &mutext)
	}

	wg.Wait()
	fmt.Println("counter=", counter)
}