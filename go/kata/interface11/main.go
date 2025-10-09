package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}

func (square Square) Area() float64 {
	return square.side * square.side
}

func (square Square) Perimeter() float64 {
	return 4 * square.side
}

func isSquare(shape Shape) bool {
	if _, ok := shape.(Square); ok {
		return true
	}

	return false
}

func main() {

	shape := Square{side: 4.56}


	if isSquare(shape) {
		fmt.Println("its a square hooray")
		return
	}

	fmt.Println("no squares detected :(")
}
