package main

import "fmt"

func main() {
	defer world() // deferred to right before main() exits
	hello()
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
