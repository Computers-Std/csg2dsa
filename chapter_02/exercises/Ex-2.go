package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 13}

	idx, steps := binarySteps(arr, 2)

	if idx != -1 {
		fmt.Printf("Found at index %d in %d steps\n", idx, steps)
	} else {
		fmt.Printf("Not found after %d steps\n", steps)
	}

}

func binarySteps(array []int, searchValue int) (int, int) {
	steps := 0

	lowerBound := 0
	upperBound := len(array) - 1

	for lowerBound < upperBound {
		steps++
		midPoint := (upperBound + lowerBound) / 2
		valueAtMidPoint := array[midPoint]

		if searchValue == valueAtMidPoint {
			return midPoint, steps
		} else if searchValue < valueAtMidPoint {
			upperBound = midPoint - 1
		} else if searchValue > valueAtMidPoint {
			lowerBound = midPoint + 1
		}
	}
	return -1, steps
}
