package main

import "fmt"

func Fibonacci(n int) int {
	return recurseFibo(n, make(map[int]int))
}

func recurseFibo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = recurseFibo(n-1, memo) + recurseFibo(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(Fibonacci(10))
}
