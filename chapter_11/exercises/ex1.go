package main

import "fmt"

// Use recursion to write a function that accepts an array of strings
// and returns the total number of characters across all the
// strings. For example, if the input array is
// ["ab", "c", "def", "ghij"]
// the output should be 10 since there are 10 characters in total.

func countCharacters(words []string) int {
	if len(words) == 0 {
		return 0
	}
	return len(words[0]) + countCharacters(words[1:])
}

func main() {
	words := []string{"ab", "c", "def", "ghij"}
	fmt.Println(countCharacters(words))
}
