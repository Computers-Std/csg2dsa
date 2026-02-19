package main

import "fmt"

func nthTraingularNum(n int) int {
	if n == 1 {
		return 1
	}
	return n + nthTraingularNum(n-1)
}

func main() {
	fmt.Println(nthTraingularNum(7))
}
