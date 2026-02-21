package main

import "fmt"

func addUntil100(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	value := nums[0] + addUntil100(nums[1:])

	if value > 100 {
		return addUntil100(nums[1:])
	} else {
		return value
	}
}

func main() {
	fmt.Println(addUntil100([]int{11, 21, 31, 41, 51, 60}))
}
