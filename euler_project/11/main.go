package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	gridText := strings.Split(string(input), "\n")

	var grid [][]int

	for i := 0; i < 20; i++ {
		row := make([]int, 20)
		grid = append(grid, row)
	}

	for i, row := range gridText {
		for j, cell := range strings.Fields(row) {
			number, err := strconv.Atoi(string(cell))

			if err != nil {
				log.Fatal(err)
			}

			grid[i][j] = number
		}
	}

	result := 0

	for y, row := range grid {
		for x, _ := range row {
			if x < 17 {
				if row[x]*row[x+1]*row[x+2]*row[x+3] > result {
					result = row[x] * row[x+1] * row[x+2] * row[x+3]
				}
			}
			if y < 17 {
				if grid[y][x]*grid[y+1][x]*grid[y+2][x]*grid[y+3][x] > result {
					result = grid[y][x] * grid[y+1][x] * grid[y+2][x] * grid[y+3][x]
				}
			}
			if y < 17 && x < 17 {
				if grid[y][x]*grid[y+1][x+1]*grid[y+2][x+2]*grid[y+3][x+3] > result {
					result = grid[y][x] * grid[y+1][x+1] * grid[y+2][x+2] * grid[y+3][x+3]
				}
			}
			if y < 17 && x > 2 {
				if grid[y][x]*grid[y+1][x-1]*grid[y+2][x-2]*grid[y+3][x-3] > result {
					result = grid[y][x] * grid[y+1][x-1] * grid[y+2][x-2] * grid[y+3][x-3]
				}
			}
		}
	}

	fmt.Println(result)
}
