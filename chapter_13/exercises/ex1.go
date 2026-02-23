package main

import "fmt"

func greatestProductOf3(nums array) int {
	length := len(nums)
	if length < 3 {
		return -1
	}
	nums.quicksort(0, length-1)
	return nums[length-1] * nums[length-2] * nums[length-3]
}

func main() {
	var arr array = []int{0, 5, 2, 1, 6, 3}
	fmt.Println(greatestProductOf3(arr))
}

type array []int

func (arr array) partion(leftPtr, rightPtr int) int {
	pivotIdx := rightPtr
	pivot := arr[pivotIdx]
	rightPtr = rightPtr - 1

	for {
		for arr[leftPtr] < pivot {
			leftPtr++
		}
		for arr[rightPtr] > pivot {
			rightPtr--
		}
		if leftPtr >= rightPtr {
			break
		} else {
			arr[leftPtr], arr[rightPtr] = arr[rightPtr], arr[leftPtr]
			leftPtr++
		}
	}
	arr[leftPtr], arr[pivotIdx] = arr[pivotIdx], arr[leftPtr]
	return leftPtr
}

func (arr array) quicksort(leftIdx, rightIdx int) {
	if (rightIdx - leftIdx) <= 0 {
		return
	}
	pivotIdx := arr.partion(leftIdx, rightIdx)
	arr.quicksort(leftIdx, pivotIdx-1)
	arr.quicksort(pivotIdx+1, rightIdx)
}
