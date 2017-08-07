package main

import "fmt"

func main() {

	// same
	//var x = 0
	//var x int

	// func expression to anonymous function
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()

	greet := makeGreeter()
	fmt.Println(greet())
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world~"
	}
}
