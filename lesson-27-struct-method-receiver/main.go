package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// (p person) is the receiver, it attaches this method to the person type
// anyone of type person can now call this method
func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"Wye Huong", "Yan", 29}
	p2 := person{first: "Wye Huong", last: "Yan", age: 29}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
