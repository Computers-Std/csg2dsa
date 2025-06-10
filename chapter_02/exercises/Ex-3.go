package main

import "fmt"

func maxBinarySteps(arraySize int) int {
	count := 0
	for arraySize > 0 {
		if arraySize%2 != 0 {
			arraySize--
			continue
		}
		arraySize = arraySize / 2
		count++
	}
	return count
}

func main() {
	fmt.Println(maxBinarySteps(100000)) // 16 steps
}
