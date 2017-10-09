package main

import (
	"fmt"
	"sort"
)

type people []string

// implements Interface interface implicity
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	// example 1
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println("before: ", studyGroup)
	sort.Sort(studyGroup)
	fmt.Println("after: ", studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("reversed: ", studyGroup)

	// example 2
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println("before: ", s)
	sort.Sort(sort.StringSlice(s))
	//sort.Strings(s)
	fmt.Println("after: ", s)
	i := sort.Reverse(sort.StringSlice(s))

	fmt.Println(i)
	fmt.Printf("%T\n", i)
	sort.Sort(i)
	fmt.Println("reversed: ", s)

	// example 3
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println("before: ", n)
	sort.Sort(sort.IntSlice(n))
	//sort.Ints(n)
	fmt.Println("after: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n))) // cannot do sort.Reverse(sort.Sort(sort.IntSlice(n))) because sort.Sort does not return an Interface
	fmt.Println("reversed: ", n)
}
