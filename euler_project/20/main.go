package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
)

func factorial(base big.Int) big.Int {
	result := big.NewInt(1)

	for i := base; i.Cmp(big.NewInt(1)) == 1; i = *i.Sub(&i, big.NewInt(1)) {
		result = result.Mul(result, &i)
	}

	return *result
}

func main() {
	factorial := factorial(*big.NewInt(100))
	factorialText := factorial.String()
	result := 0

	for _, digitText := range factorialText {
		digit, err := strconv.Atoi(string(digitText))

		if err != nil {
			log.Fatal(err)
		}

		result += digit
	}

	fmt.Println(result)
}
