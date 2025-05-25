package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

func info(g geometry) {
	fmt.Println("Geometry: ", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}

func isCircle(g geometry) {
	if c, ok := g.(*circle); ok {
		fmt.Printf("%v is a circle\n", c)
	} else {
		fmt.Printf("%v is not a circle\n", g)
	}
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r *rectangle) area() float64 {
	return r.height * r.width
}

func (r *rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	c := circle{radius: 7.65}
	r := rectangle{height: 2, width: 2.534}

	info(&c)
	info(&r)
	isCircle(&c)
	isCircle(&r)
}
