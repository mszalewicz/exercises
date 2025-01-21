package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 20 {
		wg.Add(1)

		go func() {
			hello(i+1)
			wg.Done()
		}()
	}

	wg.Wait()

}

func hello(index int) {
	rndNumber := rand.IntN(500) + 20
	time.Sleep(time.Duration(rndNumber) * time.Millisecond)

	fmt.Println("i am routines number:", index)
}
