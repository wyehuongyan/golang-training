package main

import "fmt"

func main() {
	fmt.Println(greet1("Jane ", "Doe"))  //arguments
	fmt.Println(greet2("Jane ", "Doe"))  //arguments
	fmt.Println(greet3("Jane ", "Doe ")) //arguments
}

// return
func greet1(fname, lname string) string { //parameters
	return fmt.Sprint(fname, lname)
}

// named return
func greet2(fname, lname string) (s string) { //parameters
	s = fmt.Sprint(fname, lname)
	return
}

// multiple returns
func greet3(fname, lname string) (string, string) { //parameters
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
