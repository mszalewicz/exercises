package main

import (
	"fmt"
	"math/big"
	"sort"
)

func CountUnique(values *[]big.Int) int {
	size := len(*values)
	lastValue := big.NewInt(-1)

	for _, value := range *values {
		if value.Cmp(lastValue) == 0 {
			size--
		}

		lastValue = &value
	}

	return size
}

func main() {

	uniqueValues := []big.Int{}

	for a := big.NewInt(2); a.Cmp(big.NewInt(100)) <= 0; a.Add(a, big.NewInt(1)) {
		for b := big.NewInt(2); b.Cmp(big.NewInt(100)) <= 0; b.Add(b, big.NewInt(1)) {
			newValue := new(big.Int).Exp(a, b, nil)
			uniqueValues = append(uniqueValues, *newValue)
		}
	}

	sort.Slice(uniqueValues, func(i, j int) bool {
		return uniqueValues[i].Cmp(&uniqueValues[j]) < 0
	})

	fmt.Println(CountUnique(&uniqueValues))
}
