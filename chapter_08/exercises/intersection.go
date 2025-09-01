package main

import "fmt"

func intersect(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	var intersectArray []int

	// Makeup the hashtable with items of arr1
	for _, i := range arr1 {
		m[i] = true
	}

	// Compare with arr2
	for _, j := range arr2 {
		if m[j] {
			intersectArray = append(intersectArray, j)
		}
	}
	return intersectArray
}

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{0, 2, 4, 6, 8}

	result := intersect(array1, array2)
	fmt.Println("Resultant array: ", result)
}
