package main

import (
	"fmt"
)

// create your own types

// a struct is an aggregate type (encapsulate into one data structure a bunch of fields)
type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Money", 18}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
}
