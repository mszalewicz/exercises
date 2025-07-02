package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func ShapeInfo(shape Shape) {
	fmt.Println("Area: ", shape.Area(), ", Perimeter: ", shape.Perimeter())
}

func DetectCircle(shape Shape) {
	if c, ok := shape.(*Circle); ok {
		fmt.Println(c, "is a circle")
	} else {
		fmt.Println(c, "is not a circle")
	}
}

type Circle struct {
	r float64
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.r * circle.r
}

func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.r
}

func main() {
	circle := Circle{r: 6.3453}
	ShapeInfo(&circle)
	DetectCircle(&circle)
}

