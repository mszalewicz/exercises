package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strings"
)

func contains(set []int, value int) bool {
	for _, x := range set {
		if x == value {
			return true
		}
	}
	return false
}

func findGalaxies(universe [][]byte) []image.Point {
	var galaxies []image.Point

	for y, row := range universe {
		for x, cell := range row {
			if cell == '#' {
				galaxies = append(galaxies, image.Point{x, y})
			}
		}
	}

	return galaxies
}

func expandUniverse(universe []string, rows []int, columns []int) [][]byte {

	y_offset := 0
	x_offset := 0
	var properUniverse [][]byte

	for i := 0; i < len(universe)+len(rows); i++ {
		newRow := make([]byte, len(universe[0])+len(columns))
		properUniverse = append(properUniverse, newRow)
	}

	for i := 0; i < len(properUniverse); i++ {
		for j := 0; j < len(properUniverse[0]); j++ {
			properUniverse[i][j] = '.'
		}
	}

mainloop:
	for y := 0; y < len(properUniverse); y++ {
		for x := 0; x < len(properUniverse[0]); x++ {
			if contains(rows, y-y_offset) {
				y_offset += 1
				y += 1
				continue mainloop
			} else if contains(columns, x-x_offset) {
				x_offset += 1
				x += 1
			} else {
				properUniverse[y][x] = universe[y-y_offset][x-x_offset]
			}
		}
		x_offset = 0
	}

	return properUniverse
}

func main() {
	input, err := os.ReadFile("input/input")

	if err != nil {
		log.Fatal(err)
	}

	universe := strings.Split(string(input), "\n")

	var (
		rowsToDuplicate   []int
		columsToDuplicate []int
	)

	// find empty rows
rowsloop:
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {
			if universe[y][x] == '#' {
				continue rowsloop
			}
		}
		rowsToDuplicate = append(rowsToDuplicate, y)
	}
	// find empty columns
columnsloop:
	for x := 0; x < len(universe[0]); x++ {
		for y := 0; y < len(universe); y++ {
			if universe[y][x] == '#' {
				continue columnsloop
			}
		}
		columsToDuplicate = append(columsToDuplicate, x)
	}

	newUniverse := expandUniverse(universe, rowsToDuplicate, columsToDuplicate)
	galaxies := findGalaxies(newUniverse)

	result := 0

	for i, _ := range galaxies {
		for j := 1; j < len(galaxies)-i; j++ {
			g1 := galaxies[i]
			g2 := galaxies[i+j]
			dx := int(math.Abs(float64(g1.X) - float64(g2.X)))
			dy := int(math.Abs(float64(g1.Y) - float64(g2.Y)))
			result += dx + dy
		}
	}

	fmt.Println(result)
}
