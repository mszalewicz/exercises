package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	result := float64(0)
	numbersText := strings.Split(string(input), "\n")

	for _, text := range numbersText {
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal(err)
		}
		result += number
	}

	fmt.Println(strconv.FormatFloat(result, 'f', 0, 64)[0:10])
}
