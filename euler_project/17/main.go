package main

import (
	"fmt"
)

func main() {
	length := map[int]int{
		0:  0, // for reminder logic
		1:  3, // one
		2:  3, // two
		3:  5, // three
		4:  4, // four
		5:  4, // five
		6:  3, // six
		7:  5, // seven
		8:  5, // eight
		9:  4, // nine
		10: 3, // ten
		11: 6, // eleven
		12: 6, // twelve
		13: 8, // thirteen
		14: 8, // fourteen
		15: 7, // fifteen
		16: 7, // sixteen
		17: 9, // seventeen
		18: 8, // eighteen
		19: 8, // nineteen
		20: 6, // twenty
		30: 6, // thirty
		40: 5, // forty
		50: 5, // fifty
		60: 5, // sixty
		70: 7, // seventy
		80: 6, // eighty
		90: 6, // ninety
	}

	result := 11 // <- because of "onethousand"

	for i := 1; i < 1000; i++ {
		hundreds := i / 100
		tens := (i % 100) / 10
		ones := (i % 100) % 10

		if hundreds != 0 {
			result += length[hundreds] + 7

			if tens != 0 || ones != 0 {
				result += 3
			}
		}

		if tens == 1 {
			result += length[i%100]
		} else {
			if tens != 0 {
				result += length[tens*10]
			}

			result += length[ones]
		}
	}

	fmt.Println(result)
}
