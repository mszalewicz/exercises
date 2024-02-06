package main

import (
	"fmt"
	"time"
)

func main() {

	currDate := time.Date(1901, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	sundayCounter := 0

	for {
		if currDate.Weekday() == time.Sunday {
			sundayCounter += 1
		}

		currDate = currDate.AddDate(0, 1, 0)
		currYear, _, _ := currDate.Date()

		if currYear > 2000 {
			break
		}
	}

	fmt.Println(sundayCounter)
}
