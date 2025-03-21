package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"runtime"
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
			for range 40 {
				fmt.Print(".")
				time.Sleep(delay)
			}
			fmt.Print("\r\033[K")
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
	var mut sync.Mutex
	jobs := make(chan string)
	done := make(chan bool)

	file, err := os.OpenFile("test", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	for range runtime.NumCPU() {
		go func(jobs chan string, file *os.File) {
			for job := range jobs {
				mut.Lock()
				_, err := file.WriteString(job + "\n")
				if err != nil {
					log.Fatal(err)
				}
				mut.Unlock()
			}
		}(jobs, file)
	}

	go spinner(85*time.Millisecond, done)

	for range 100_000 {
		rndString, err := randomUTF8String(500)
		if err != nil {
			log.Fatal(err)
		}

		jobs <- rndString
	}

	close(jobs)
	done <- true

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}
}
