package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct { // outer type
	Person        // embedded type (inner type)
	First         string
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("Hello!")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Ni Hao!")
}

func main() {
	// fields and methods of inner type Person is promoted to outer type DoubleZero
	p1 := DoubleZero{
		Person: Person{
			First: "Wye Huong",
			Last:  "Yan",
			Age:   29,
		},
		First:         "OOP", // outer type field overwrites inner type field
		LicenseToKill: false,
	}

	p2 := &Person{"A", "B", 1}

	// overriding fields
	fmt.Println(p1.First)
	fmt.Println(p1.Person.First) // but we still can access the inner type's First
	fmt.Println(p1.LicenseToKill)

	// overriding methods
	p1.Greeting()        // outer
	p1.Person.Greeting() // inner

	// struct pointer
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
	fmt.Println(p2.First) // *p2.First behind the scenes
}
