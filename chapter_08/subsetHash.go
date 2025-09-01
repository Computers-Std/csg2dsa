package main

import "fmt"

func isSubset(arr1, arr2 []int) bool {
	var largeArray, smallArray []int
	m := make(map[int]bool)

	if len(arr1) > len(arr2) {
		largeArray = arr1
		smallArray = arr2
	} else {
		largeArray = arr2
		smallArray = arr1
	}

	// Store values of the larger array in hash table
	for _, key := range largeArray {
		m[key] = true
	}

	// Check if all values of the smaller array exist in the hash table
	for _, i := range smallArray {
		if !m[i] {
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
