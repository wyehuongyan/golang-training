package main

import (
	"fmt"
)

func main() {
	fmt.Println(x)
	foo()
	y := 10 // syntactic sugar for var y int = 10
	fmt.Println(y)
	foo3()
}

func foo() {
	fmt.Println(x)
}

func foo2() {
	x := 0

	// func expression with anonymous function
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func foo3() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

var x = 42
