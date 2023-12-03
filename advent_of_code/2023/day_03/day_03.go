package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Gear struct {
	ratio           int
	adjacentNumbers int
}

type Point struct {
	x, y int
}

func checkNeighbours(rows *[]string, x, y, max_height, max_length *int) (bool, []Point) {

	hasNeighbours := false
	var neighbourGearsCoordinates []Point

	for i := -1; i < 2; i++ {
		if *y+i < 0 || *y+i == *max_height {
			continue
		}
		for j := -1; j < 2; j++ {
			if (*x+j < 0 || *x+j == *max_length) || (i == 0 && j == 0) || (i == 0 && unicode.IsDigit(rune((*rows)[*y+i][*x+j]))) {
				continue
			}
			if (*rows)[*y+i][*x+j] != '.' {
				hasNeighbours = true
				if (*rows)[*y+i][*x+j] == '*' {
					neighbourGearsCoordinates = append(neighbourGearsCoordinates, Point{*x + j, *y + i})
				}
			}
		}
	}
	return hasNeighbours, neighbourGearsCoordinates
}

func populateGearInfo(coordinates *[]Point, gears *[][]Gear, value *int) {

	var uniquePoints []Point

	for _, point := range *coordinates {
		if !slices.Contains(uniquePoints, point) {
			uniquePoints = append(uniquePoints, point)
		}
	}

	for _, point := range uniquePoints {
		(*gears)[point.x][point.y].adjacentNumbers += 1
		(*gears)[point.x][point.y].ratio *= *value
	}
}

func main() {

	start := time.Now()
	input, _ := os.ReadFile("inputs/day_03.txt")
	rows := strings.Split(string(input), "\n")
	result1 := 0
	max_height := len(rows)
	max_length := len((rows)[0])
	var number []rune

	gears := make([][]Gear, max_height)

	for z := 0; z < max_height; z++ {
		gears_row := make([]Gear, max_length)
		for k := 0; k < max_length; k++ {
			gears_row[k] = Gear{ratio: 1, adjacentNumbers: 0}
		}
		gears[z] = gears_row
	}

	/*
		Original idea of creating [][]Gear, duplicating the same []Gear
		row over multiple rows. Lead to situation were given y poisiton was
		updated the same over all rows.

			gears_row := make([]Gear, max_length)
			for k := 0; k < max_length; k++ {
				gears_row[k] = Gear{ratio: 1, adjacentNumbers: 0}
			}
			for z := 0; z < max_height; z++ {
				gears[z] = gears_row
			}
	*/

	for y, row := range rows {
		hasNeighbours := false
		var gearPoints []Point
		for x, char := range row {
			if unicode.IsDigit(char) {
				number = append(number, char)
				if !hasNeighbours {
					hasNeighbours, gearPoints = checkNeighbours(&rows, &x, &y, &max_height, &max_length)
				}
			} else if len(number) != 0 && hasNeighbours {
				value, _ := strconv.Atoi(string(number))
				result1 += value
				hasNeighbours = false
				number = nil
				if len(gearPoints) != 0 {
					populateGearInfo(&gearPoints, &gears, &value)
				}
			} else {
				hasNeighbours = false
				number = nil
			}
		}
		if hasNeighbours {
			value, _ := strconv.Atoi(string(number))
			result1 += value
			hasNeighbours = false
			number = nil
			if len(gearPoints) != 0 {
				populateGearInfo(&gearPoints, &gears, &value)
			}
		}
	}

	result2 := 0

	for _, row := range gears {
		for _, gear := range row {
			if gear.adjacentNumbers == 2 {
				result2 += gear.ratio
			}
		}
	}

	fmt.Println(result1)
	fmt.Println(result2)

	elapsed := time.Since(start)
	fmt.Println("time " + fmt.Sprint(elapsed))
}
