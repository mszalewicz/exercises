package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Number of possible lattice routesfrom (0, 0) to (n, k),
	// can be obtained from binomial coefficient (n+k choose n).
	// Here lattice is square, hence formula looks like (2n choose n).
	//
	// Furthermore, formula can be calculated as (in greek pi notation):
	// (i = 1 to n) -> (n+i) / i

	var n float64 = 20
	var result float64 = 1

	for i := float64(1); i <= n; i++ {
		result *= (n + i) / i
	}

	fmt.Println(strconv.FormatFloat(result, 'f', 0, 64))
}
