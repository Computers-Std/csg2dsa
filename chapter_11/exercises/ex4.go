package main

import "fmt"

func indexOfX(word string, currIdx int) int {
	if word[0] == 'x' {
		return currIdx
	}
	return indexOfX(word[1:], currIdx+1)
}

func indexOfXv2(word string) int {
	if word[0] == 'x' {
		return 0
	}
	return indexOfXv2(word[1:]) + 1
}

func main() {
	word := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(indexOfX(word, 0))
	fmt.Println(indexOfXv2(word))
}
