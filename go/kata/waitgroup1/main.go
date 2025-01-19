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
			printInfo(i)
			defer wg.Done()
		}()
	}

	wg.Wait()
}

func printInfo(index int) {
	randInterval := rand.IntN(3)+1

	time.Sleep(time.Duration(randInterval) * time.Second)
	fmt.Println("Goroutine:", index+1)
}
