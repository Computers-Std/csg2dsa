package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	return recurseFibo(n, memo)
}

func recurseFibo(n int, memo []int) int {
	if n <= 1 {
		return n
	}
	if memo[n] != 0 { // already calculated
		return memo[n]
	}
	memo[n] = recurseFibo(n-1, memo) + recurseFibo(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(Fibonacci(10))
}
