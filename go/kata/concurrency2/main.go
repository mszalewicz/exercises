package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"
	"time"
)

func spinner(delay time.Duration, done chan bool) {
	for {
		select {
		case <-done:
			fmt.Print("\r\033[K")
			return
		default:

			for {
				for range 40 {
					fmt.Print(".")
					time.Sleep(delay)
				}
				fmt.Print("\r\033[K")
			}
		}
	}
}

func randomUTF8String(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?/~"
	result := make([]rune, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = rune(charset[num.Int64()])
	}

	return string(result), nil
}

func main() {
	done := make(chan bool)
	go spinner(70*time.Millisecond, done)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for range 100000 {
		go func(mutex *sync.Mutex) {
			wg.Add(1)

			randText, err := randomUTF8String(500)
			randText = randText + "\n"

			if err != nil {
				log.Fatal(err)
			}

			mutex.Lock()

			file, err := os.OpenFile("testFile", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			file.WriteString(randText)

			err = file.Close()

			if err != nil {
				log.Fatal(err)
			}

			mutex.Unlock()
			wg.Done()
		}(&mutex)
	}

	wg.Wait()
	done <- true
}
