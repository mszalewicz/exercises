package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func findStart(layout *[]string) Point {

	var x, y int

loops:
	for j, row := range *layout {
		for i, value := range row {
			if value == 'S' {
				x, y = i, j
				break loops
			}
		}
	}

	return Point{x, y}
}

func traverse(x, y int, layout *[]string) (int, int) {
	var starters []Point
	var vertices []Point

	steps := 1
	area := 0

	up := (*layout)[y-1][x]
	down := (*layout)[y+1][x]
	left := (*layout)[y][x-1]
	right := (*layout)[y][x+1]

	if strings.ContainsRune("|7F", rune(up)) {
		starters = append(starters, Point{x, y - 1})
	}
	if strings.ContainsRune("|JL", rune(down)) {
		starters = append(starters, Point{x, y + 1})
	}
	if strings.ContainsRune("-LF", rune(left)) {
		starters = append(starters, Point{x - 1, y})
	}
	if strings.ContainsRune("-7J", rune(right)) {
		starters = append(starters, Point{x + 1, y})
	}

	currNode := Point{x, y}
	nextNode := starters[0]

	vertices = append(vertices, currNode)

	for {
		switch (*layout)[nextNode.y][nextNode.x] {
		case '-':
			if nextNode.x-currNode.x > 0 {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.x += 1
			} else {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.x += -1
			}
		case '|':
			if nextNode.y-currNode.y > 0 {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.y += 1
			} else {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.y += -1
			}
		case 'F':
			if nextNode.y-currNode.y == 0 {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.y += 1
			} else {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.x += 1
			}
			vertices = append(vertices, currNode)
		case 'L':
			if nextNode.y-currNode.y == 0 {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.y += -1
			} else {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.x += 1
			}
			vertices = append(vertices, currNode)
		case 'J':
			if nextNode.y-currNode.y == 0 {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.y += -1
			} else {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.x += -1
			}
			vertices = append(vertices, currNode)
		case '7':
			if nextNode.y-currNode.y == 0 {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.y += 1
			} else {
				currNode.x = nextNode.x
				currNode.y = nextNode.y
				nextNode.x += -1
			}
			vertices = append(vertices, currNode)
		}

		steps += 1
		if (*layout)[nextNode.y][nextNode.x] == 'S' {
			break
		}

	}

	// Shoelace formula - https://en.wikipedia.org/wiki/Shoelace_formula
	for i := 0; i < len(vertices); i++ {
		x1 := vertices[i].x
		y1 := vertices[i].y

		x2 := vertices[(i+1)%len(vertices)].x
		y2 := vertices[(i+1)%len(vertices)].y

		area += x1*y2 - x2*y1
	}

	// Pick's theorem - https://en.wikipedia.org/wiki/Pick%27s_theorem
	return steps / 2, int(math.Abs(float64(area)))/2 + 1 - (steps / 2)
}

func main() {

	input, err := os.ReadFile("input/day_10")

	if err != nil {
		log.Fatal(err)
	}

	layout := strings.Split(string(input), "\n")
	staringPoint := findStart(&layout)
	furthestPoint, area := traverse(staringPoint.x, staringPoint.y, &layout)

	fmt.Println(furthestPoint)
	fmt.Println(area)
}
