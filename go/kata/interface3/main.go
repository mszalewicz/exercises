package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
	area() float64
}

func Information(shape Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", shape.area(), shape.perimeter())
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func isCircle(shape Shape) {
	if c, ok := shape.(Circle); ok {
		fmt.Println("It is circle.")
		Information(c)
	} else {
		fmt.Println("It is not circle.")
	}
}

func main( ) {
	circle := Circle{radius: 65.8}

	isCircle(circle)
}