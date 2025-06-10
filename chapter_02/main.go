package main

import "fmt"

func main() {
	arr := []int{3, 17, 75, 80, 202}

	// result := linear_search(arr, 22)
	result := binarySearch(arr, 22)

	if result != -1 {
		fmt.Println("Found at index: ", result)
	} else {
		fmt.Println("Not found")
	}

}

func linearSearch(array []int, searchValue int) int {
	// we iterate through every elemeny in the array:
	for index, element := range array {
		// If we find the value we're looking for, we return its index
		if element == searchValue {
			return index
		}
		// If we reach an element that is greater than the value we're
		// looking for, we can exit the loop early
		if element > searchValue {
			break
		}
	}
	return -1 // indicates "not found"
}

func binarySearch(array []int, searchValue int) int {
	/* First, we establish the lower and upper bounds of where the
	   value we're searching for can be. To start, the lower bound is the
	   first value in the array, while the upper bound is the last value */

	lowerBound := 0
	upperBound := len(array) - 1

	/* we begin a loop in which we keep inspecting the middlemost
	   value between the upper and lower bounds */

	for lowerBound <= upperBound {
		midPoint := (upperBound + lowerBound) / 2
		valueAtMidPoint := array[midPoint]

		if searchValue == valueAtMidPoint {
			return midPoint
		} else if searchValue < valueAtMidPoint {
			upperBound = midPoint - 1
		} else if searchValue > valueAtMidPoint {
			lowerBound = midPoint + 1
		}
	}
	return -1 // not found
}
