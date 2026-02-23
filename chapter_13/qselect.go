package main

import "fmt"

func (arr sortableArray) quickselect(kthLow, leftIdx, rightIdx int) int {
	// If we reach a base case - that is, that the subarray has one
	// cell, we know we've found the value we're looking for:
	if (rightIdx - leftIdx) <= 0 {
		return arr[leftIdx]
	}

	// Partition the array and grab the index of pivot
	pivotIdx := arr.partition(leftIdx, rightIdx)

	// If what we're looking for is to the left of the pivot
	if kthLow < pivotIdx {
		// Recursively perform quickselect on the subarray to the left
		// of the pivot.
		return arr.quickselect(kthLow, leftIdx, pivotIdx-1)
	} else if kthLow > pivotIdx {
		// If what we're looking for is to the right of the pivot

		// Recursively perform quickselect on the subarray to the
		// right of the pivot
		return arr.quickselect(kthLow, pivotIdx+1, rightIdx)
	} else { // kthLow == pivotIdx
		// If after the partition, the pivot position is in the same
		// spot as the kth lowest value, we've found the value we're
		// looking for
		return arr[pivotIdx]
	}
}

func main() {
	var array sortableArray = []int{0, 5, 2, 1, 6, 3}
	val := array.quickselect(4, 0, len(array)-1) // 5
	fmt.Println(val)
}
