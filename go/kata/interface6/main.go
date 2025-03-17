package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
	area() float64
}

func info(shape Shape) {
	fmt.Println("Perimeter: ", shape.perimeter())
	fmt.Println("Area: ", shape.area())
}

type Circle struct {
	radius float64
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func isCircle(shape Shape) {
	if val, ok := shape.(Circle); ok {
		fmt.Println(val, " is indeed a circle")
	} else {
		fmt.Println(val, " is not a circle")
	}
}

type Square struct{}

func (squate Square) perimeter() float64 {
	return 0
}

func (squate Square) area() float64 {
	return 0
}

func main() {
	circle := Circle{radius: 4.567}
	square := Square{}

	info(circle)
	isCircle(circle)
	isCircle(square)
}
