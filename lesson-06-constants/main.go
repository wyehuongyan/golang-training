package main

import "fmt"

const p string = "death & taxes"
const q = 31

const (
	A = iota
	B = iota
	C = iota
)

func main() {
	const q = 2333

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
