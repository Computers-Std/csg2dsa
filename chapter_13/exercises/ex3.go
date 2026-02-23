package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(ONGreatest([]int{1, 2, 3, 4, 5}))
	fmt.Println(ONlogNGreatest([]int{1, 2, 3, 4, 5}))
	fmt.Println(ONSqGreatest([]int{1, 2, 3, 4, 5}))
}

func ONGreatest(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	n := nums[0]
	for _, val := range nums {
		if val > n {
			n = val
		}
	}
	return n
}

func ONlogNGreatest(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	slices.Sort(nums)
	return nums[len(nums)-1]
}

func ONSqGreatest(nums []int) int {
	for _, i := range nums {
		isGreatest := true
		for _, j := range nums {
			if j > i {
				isGreatest = false
				break
			}
		}
		if isGreatest {
			return i
		}
	}
	return -1
}
