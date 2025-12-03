package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		tens := line[0]
		ones := line[1]

		for i := 1; i < len(line)-1; i++ {
			if line[i] > tens {
				tens = line[i]
				ones = line[i+1]
				continue
			}

			if line[i] > ones {
				ones = line[i]
			}
		}

		if line[len(line)-1] > ones {
			ones = line[len(line)-1]
		}

		counter += int(tens - '0')*10 + int(ones - '0')
	}

	fmt.Println(counter)
}
