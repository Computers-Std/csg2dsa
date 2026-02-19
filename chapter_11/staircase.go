package main

import "fmt"

func numOfPath(n int) int {
	switch {
	case n < 0:
		return 0
	case n == 0, n == 1:
		return 1
	}
	return numOfPath(n-1) + numOfPath(n-2) + numOfPath(n-3)
}

func main() {
	fmt.Println(numOfPath(3))
}
