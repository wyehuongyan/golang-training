package main

import (
	"fmt"
)

func main() {
	// decimal binary
	fmt.Printf("%d - %b\n", 42, 42)

	// hexadecimal
	fmt.Printf("%#x\n", 42)
	fmt.Printf("%#X\n", 42)

	// loops
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %X \n", i, i, i)
	}

	// utf-8 loop
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %X \t %q \n", i, i, i, i)
	}
}
