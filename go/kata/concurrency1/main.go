package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"sync"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?/~"

func randomUTF8String(length int) (string, error) {
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
	var waitgroup sync.WaitGroup
	var mutex sync.Mutex

	for range 1000 {
		go func() {
			waitgroup.Add(1)

			randText, err := randomUTF8String(16)

			if err != nil {
				log.Fatal(err)
			}

			mutex.Lock()

			file, err := os.OpenFile("testFile.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			if err != nil {
				log.Fatal(err)
			}

			_, err = file.WriteString(randText)

			if err != nil {
				log.Fatal(err)
			}

			err = file.Close()

			if err != nil {
				log.Fatal(err)
			}

			mutex.Unlock()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
}
