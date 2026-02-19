package main

import "fmt"

func anagramsOf(word string) []string {
	if len(word) <= 1 {
		return []string{word}
	}
	collection := []string{}

	first := word[0]
	rest := word[1:]

	subAnagrams := anagramsOf(rest) // recursion

	for _, subAn := range subAnagrams {
		// insert 'first' in every possible position
		for i := 0; i <= len(subAn); i++ {
			res := subAn[:i] + string(first) + subAn[i:]
			collection = append(collection, res)
		}
	}
	return collection
}

func main() {
	result := anagramsOf("bolt")
	fmt.Println(result)
}

// Efficiency
// O(N!)
