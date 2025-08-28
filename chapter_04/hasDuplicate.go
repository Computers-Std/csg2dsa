package main

import "fmt"

func main() {
	numbers := []int{1, 4, 5, 2, 9}
	steps, hasDup := n2HasDuplicate(numbers)
	fmt.Printf("Steps: %d, Has duplicate: %v\n", steps, hasDup)
}

func n2HasDuplicate(nums []int) (int, bool) {
	steps := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			steps++
			if i != j && nums[i] == nums[j] {
				return steps, true
			}
		}
	}
	return steps, false
}

func linearHasDuplicate(nums []int) (int, bool) {
	existingNumbers := []int
	for i := 0; i < len(nums); i++ {
		if existingNumbers[nums[i]] == 1 {
			return true
		} else {
			existingNumbers[nums[i]] = 1
		}
	}
	return false
}
