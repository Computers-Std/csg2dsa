package main

import "fmt"

func missingAlpha(arr string) rune {
	m := make(map[rune]bool)

	for _, i := range arr {
		m[i] = true
	}

	for i := 'a'; i <= 'z'; i++ {
		if !m[i] {
			return i
		}
	}
	return 0
}

func main() {
	strArr := "the quick brown box jumps over a lazy dog"
	res := missingAlpha(strArr)

	if res == 0 {
		fmt.Println("No missing letters")
	} else {
		fmt.Printf("Missing letter is %c\n", res)
	}
}
