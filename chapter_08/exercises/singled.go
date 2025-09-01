package main

import "fmt"

func nonDuplicate(arr string) rune {
	m := make(map[rune]int)

	for _, r := range arr {
		m[r]++
	}

	for _, c := range arr {
		if m[c] == 1 {
			return c
		}
	}
	return 0
}

func main() {
	str := "minimum"
	res := nonDuplicate(str)
	fmt.Printf("Non-Duplicate characeter is %c", res)
}
