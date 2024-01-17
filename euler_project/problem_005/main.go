package main

import (
	"fmt"

	"github.com/mszalewicz/simplemath"
)

func main() {

	var set []int

	for i := 1; i <= 20; i++ {
		set = append(set, i)
	}

	// Smallest multiple implemented in my package simplemath
	// For details visit github.com/mszalewicz/simplemath
	result := simplemath.LCM(set[0], set[1], set[2:]...)

	fmt.Println(result)
}
