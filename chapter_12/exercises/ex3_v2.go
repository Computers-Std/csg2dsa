package main

import "fmt"

func uniqPaths(rows, cols int) int {
	// initialize the cache 2d-array
	cache := make([][]int, rows+1)
	for i := range cache {
		cache[i] = make([]int, cols+1)
	}
	return recurseUniqPaths(rows, cols, cache)
}

func recurseUniqPaths(rows, cols int, cache [][]int) int {
	if rows == 1 || cols == 1 {
		return 1
	}

	if cache[rows][cols] != 0 {
		return cache[rows][cols]
	}

	cache[rows][cols] = recurseUniqPaths(rows-1, cols, cache) +
		recurseUniqPaths(rows, cols-1, cache)

	return cache[rows][cols]
}

func main() {
	fmt.Println(uniqPaths(3, 7))
}
