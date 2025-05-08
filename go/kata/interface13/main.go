package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func Info(shape Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", shape.area(), shape.perimeter())
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return circle.radius * circle.radius * math.Pi
}

func (circle Circle) perimeter() float64 {
	return 2 * circle.radius * math.Pi
}

func isCircle(shape Shape) {
	if _, valid := shape.(Circle); valid {
		fmt.Println("it is a circle")
	} else {
		fmt.Println("not a circle")
	}
}

func main() {
	c := Circle{radius: 895.67}

	Info(c)
	isCircle(c)
}
