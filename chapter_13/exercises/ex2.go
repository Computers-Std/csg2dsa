package main

import (
	"fmt"
	"slices"
)

func missingNum(nums []int) int {
	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 6}
	fmt.Println(missingNum(nums))
}
