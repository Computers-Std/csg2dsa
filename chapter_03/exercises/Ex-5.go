package main

import (
	"fmt"
)

func median(array []int) float64 {
	length := len(array)
	middle := length / 2

	if length%2 == 0 {
		return float64(array[middle-1]+array[middle]) / 2
	}
	return float64(array[middle])
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 9, 8, 7, 6}
	result := median(numbers)
	fmt.Printf("Median: %.1f\n", result)
}

// O(1), No matter the number of items in the array, every time It
// will just perform one calculation to find the median of the array
