package main

import "fmt"

type pair struct {
	rows int
	cols int
}

func uniqPaths(rows, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return recurseUniqPaths(rows, cols, make(map[pair]int))
}

func recurseUniqPaths(rows, cols int, memo map[pair]int) int {
	r, c := rows, cols
	if r > c {
		r, c = c, r
	}
	curr := pair{r, c}

	if val, ok := memo[curr]; ok {
		return val
	}
	if rows == 1 || cols == 1 {
		return 1
	}

	result := recurseUniqPaths(rows-1, cols, memo) +
		recurseUniqPaths(rows, cols-1, memo)

	memo[curr] = result

	return result
}

func main() {
	fmt.Println(uniqPaths(3, 7))
}
