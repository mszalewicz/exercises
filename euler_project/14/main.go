package main

import "fmt"

func main() {
	chainCount := 0
	chainStartingNumber := 0

	for i := 2; i < 1000000; i++ {
		count := 0
		for j := i; j != 1; {
			count += 1
			if j%2 == 0 {
				j = j / 2
			} else {
				j = 3*j + 1
			}
		}
		if count > chainCount {
			chainCount = count
			chainStartingNumber = i
		}
	}

	fmt.Println(chainStartingNumber)
}
