package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		letter := 'A'
		fmt.Println(letter)
		fmt.Printf("%T \n", letter)

		character := rune("A"[0]) // a string is a slice of bytes (collection of runes)
		fmt.Println(character)

		word := "hello"
		letter = rune(word[0])
		fmt.Println(letter)

		// integer to string letter
		for i := 65; i <= 122; i++ {
			fmt.Println(i, " - ", string(i), " - ", i%12)
		}

		// hashbucket example
		n := HashBucket("Go", 12)
		fmt.Println(n)
	*/

	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 12)

	// loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

// HashBucket : oh yeah
func HashBucket(word string, numBuckets int) int {
	letter := int(word[0])
	bucket := letter % numBuckets

	return bucket
}
