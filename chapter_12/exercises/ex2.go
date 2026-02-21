package main

import "fmt"

func Golomb(n int) int {
	if n == 1 {
		return 1
	}
	return recurseGolomb(n, make(map[int]int))
	// return 1 + Golomb(n-Golomb(Golomb(n-1)))
}

func recurseGolomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = recurseGolomb(n-recurseGolomb(
		recurseGolomb(n-1, memo), memo), memo) + 1

	return memo[n]
}

func main() {
	fmt.Println(Golomb(15))
}
