package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57) // just a bunch of comma separated arguments
	fmt.Println(n)

	data := []float64{43, 56, 87, 12, 45, 57} // a slice of float 64
	m := average(data...)                     // need to use ... to pull each item out to pass it in (variadic arguments)
	fmt.Println(m)
}

// accepts unlimited number of float64 numbers as parameters (variadic parameters)
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	total := 0.0
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
