package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	base := math.Pow(2., 1000.)
	baseText := strconv.FormatFloat(base, 'f', 0, 64)
	result := 0

	for _, text := range baseText {
		digit, err := strconv.Atoi(string(text))

		if err != nil {
			log.Fatal(err)
		}

		result += digit
	}

	fmt.Println(result)
}
