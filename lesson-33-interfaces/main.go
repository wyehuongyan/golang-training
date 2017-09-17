package main

import (
	"fmt"
	"math"
)

// first shape
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

// second shape
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Shape interface
// the type Shape is an interface (defines functionality)
// anything that has this method signature, implicitly implements this interface
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area()) // polymorphism
}

func main() {
	s := Square{10}
	c := Circle{5}

	info(s)
	info(c)
}
