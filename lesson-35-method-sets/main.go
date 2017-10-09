package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

type Shape interface {
	area() float64
}

// method-set with value receiver
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// method-set with pointer receiver
func (s *Square) area() float64 {
	return s.side * s.side
}

func info(s Shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	c := Circle{10}
	s := Square{10}

	// value receivers can accept both value and pointer value
	info(c)
	info(&c)

	// pointer receiver can only accept pointer value
	info(&s)
}

// Receivers	Values
// ----------------------------
// (t T)		T and *T
// (t *T)		*T
