package main

import (
	"fmt"
	"image"
)

type Object struct {
	id       string
	position image.Point
}

func (object Object) String() string {
	return fmt.Sprintf("ID: %s, x: %d, y: %d", object.id, object.position.X, object.position.Y)
}

func main() {
	object := Object{id: "dfg924", position: image.Point{X: 30, Y: 89}}

	fmt.Println(object)
}
