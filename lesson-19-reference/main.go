package main

import "fmt"

func main() {
	m := make([]string, 1, 25)

	fmt.Println(m)

	changeMeArray(m)
	fmt.Println(m)

	n := make(map[string]int)
	changeMeMap(n)
	fmt.Println(n["Todd"])

	a := 13

	changeMeValue(&a)
	fmt.Println(a)
}

// pass by value (slice of string)
// a slice is a reference type (an address pointing to an underlaying array)
func changeMeArray(z []string) {
	z[0] = "Todd"
	fmt.Println(z)
}

func changeMeMap(z map[string]int) {
	z["Todd"] = 44
}

func changeMeValue(b *int) {
	*b = 15
}
