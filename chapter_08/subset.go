package main

import "fmt"

func isSubset[T comparable](arr1, arr2 []T) bool {
	var largeArray, smallArray []T
	if len(arr1) > len(arr2) {
		largeArray = arr1
		smallArray = arr2
	} else {
		largeArray = arr2
		smallArray = arr1
	}

	for i := 0; i < len(smallArray); i++ {
		foundMatch := false

		for j := 0; j < len(largeArray); j++ {
			if smallArray[i] == largeArray[j] {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			return false
		}
	}
	return true
}

func main() {
	arrLarge := []int{1, 3, 4, 5, 6}
	arrSmall := []int{3, 4}

	res := isSubset(arrLarge, arrSmall)
	fmt.Println(res)
}
