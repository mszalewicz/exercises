package main

import "fmt"

func main() {

	a := 1
	b := 1
	tmp := 0
	result := 0

	for {

		if b > 4000000 {
			break
		}

		tmp = a + b

		if tmp%2 == 0 {
			result += tmp
		}

		a = b
		b = tmp
	}

	fmt.Println(result)
}
