package main

import "fmt"

func main() {
	switch "allah" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi", "allah":
		fmt.Println("Wassup Medhi")
	default:
		fmt.Println("default")
	}
}
