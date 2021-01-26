// Follow the comments below...

package main

import (
	"fmt"
	"math"
)

// Create a type of "square" with underlying type "struct"
type square struct {
	length float64
}

// Create a type of "circle" with underlying type "struct"
type circle struct {
	radius float64
}

// Attach a method to each the calculates the area of each shape and returns it

// Square area method
func (square1 square) calculateArea() float64 {
	area := square1.length * square1.length
	return area
}

// Circle area method
func (circle1 circle) calculateArea() float64 {
	area := math.Pi * (circle1.radius * circle1.radius)
	return area
}

// Create a type of "shape" which defines an interface as anything that has the area method
type shape interface {
	calculateArea() float64 // This reflects the signature of the above functions
}

// Create a func "info" which takes type "shape" and then prints the area
func info(shape1 shape) {
	fmt.Printf("The area of the shape provided is: %v\n", shape1.calculateArea())
}

func main() {

	// Create a VALUE of TYPE square
	square1 := square{
		length: 4,
	}

	// Create a VALUE of TYPE circle
	circle1 := circle{
		radius: 8,
	}

	// Use func info to print the area of the square
	info(square1)

	// Use func info to print the area of the circle
	info(circle1)
}
