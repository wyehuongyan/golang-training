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

	greeting() // Hello world!

	greet := makeGreeter()
	fmt.Println(greet()) // Hello world~

	// anonymous self executing function
	func() {
		fmt.Println("I'm driving!")
	}()
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world~"
	}
}
