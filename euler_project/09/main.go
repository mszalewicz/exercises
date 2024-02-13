package main

import "fmt"

func main() {
	result := 0
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			for c := b + 1; c < 1000; c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					result = a * b * c
				}
			}
		}
	}
	fmt.Println(result)
}
