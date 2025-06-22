package main

import (
	"fmt"
	"strings"
)

func selectAStrings(words []string) []string {
	var Awords []string

	for _, word := range words {
		if strings.HasPrefix(word, "a") {
			Awords = append(Awords, word)
		}
	}
	return Awords
}

func main() {
	words := []string{"apple", "banana", "avocado", "grape"}
	matches := selectAStrings(words)
	fmt.Println("Words starting with 'a': ", matches)
}

// O(N), all the items in the Array are checked in a linear fashion
