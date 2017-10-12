package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // b is a pointer to the memory address of a

	// *int is a integer pointer (points to a mem address which an int is stored)

	fmt.Println(b)
	fmt.Println(*b) // dereferencing a pointer to get the value (* is an operator)

	*b = 42 // change the value at this address into 42
	fmt.Println(a)
	fmt.Println(*b)

	/* everything in go is passed by value!! */
	/* the example above is passed by pointer (&a and b are two pointers pointing to the value 43 stored at the same mem location) */

	x := 5

	zero1(x) // only passing the value

	fmt.Println("value of x was not changed:", x)

	fmt.Printf("same mem address %p \n", &x)

	zero2(&x) // passing in mem address

	fmt.Println("value of x was changed:", x)
}

func zero1(x int) {
	x = 0
}

func zero2(x *int) {
	*x = 0 // dereference the pointer to get value and assign it to 0

	fmt.Printf("same mem address %p \n", x)
}
