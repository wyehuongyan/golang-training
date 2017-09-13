package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11} // composite literal
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	capTest() // capacity increased
	sliceAppendSlice()
	createMultiSlice()
	createSlices()
	incrementSliceValue()
}

func incrementSliceValue() {
	mySlice := make([]int, 1, 1)

	fmt.Println(mySlice)
	mySlice[0] = 8

	fmt.Println(mySlice)
	mySlice[0]++

	fmt.Println(mySlice)
}

func createSlices() {
	// shorthand
	student1 := []string{} // with empty {} literal (no values)
	students1 := [][]string{}

	fmt.Println(student1)
	fmt.Println(students1)
	fmt.Println(student1 == nil)

	// var
	var student2 []string // will be set to zero value of slice, which is nil (no underlying slice created yet)
	var students2 [][]string

	fmt.Println(student2)
	fmt.Println(students2)
	fmt.Println(student2 == nil)

	// make
	student3 := make([]string, 35)
	students3 := make([][]string, 35)

	fmt.Println(student3)
	fmt.Println(students3)
	fmt.Println(student3 == nil)
}

func createMultiSlice() {
	multiSlice := make([][]string, 0)

	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	// store
	multiSlice = append(multiSlice, student1)

	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	// store
	multiSlice = append(multiSlice, student2)

	// print
	fmt.Println(multiSlice)
}

func capTest() {
	mySlice := make([]int, 0, 5)

	fmt.Println("------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}

func sliceAppendSlice() {
	// appending slice to slice
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	weekends := []string{"Saturday", "Sunday"}

	weekdays = append(weekdays, weekends...)

	fmt.Println(weekdays)
}
