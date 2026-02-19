package main

import "fmt"

func uniqPaths(rows, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return uniqPaths(rows-1, cols) + uniqPaths(rows, cols-1)
}

func main() {
	fmt.Println(uniqPaths(3, 7))
}
