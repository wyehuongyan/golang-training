package main

import (
	"fmt"
	"strconv"
)

/*
// MY SOLUTION
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 0 // init c1 by sending 0 to c1
	}()

	// 1 reads from c1 and writes into c2
	// 2 reads from c2 and writes into c1
	for i := 0; i < 20; i++ {
		incrementor("1", c1, c2)
		incrementor("2", c2, c1)
	}

	fmt.Println("c1: ", <-c1)
}

func incrementor(s string, read chan int, write chan int) {
	go func() {
		k := <-read
		write <- k + 1

		fmt.Printf("Incrementor: %s Reading: %d Writing %d\n", s, k, k+1)
	}()
}
*/

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/

func main() {

	c := incrementor(2)

	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}
