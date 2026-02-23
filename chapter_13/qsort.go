package main

import "fmt"

func (arr sortableArray) quicksort(leftIdx, rightIdx int) {
	// Base case: the subarray hsa 0 or 1 elements
	if (rightIdx - leftIdx) <= 0 {
		return
	}

	// Partition the range of elements and grab the index of the pivot
	pivotIdx := arr.partition(leftIdx, rightIdx)

	// Recursively call the quicksort method on whatever is to the left of
	// the pivot.
	arr.quicksort(leftIdx, pivotIdx-1)
	// Recursively call the quicksort method on whatever is to the right of
	// the pivot.
	arr.quicksort(pivotIdx+1, rightIdx)
}

func main() {
	var array sortableArray = []int{0, 5, 2, 1, 6, 3}
	array.quicksort(0, len(array)-1)
	fmt.Println(array)
}

// Efficiency of Quicksort:
// For Partition: O(N)
// On Average: O(N LogN)
