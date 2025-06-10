package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 13}

	idx, steps := linearSteps(arr, 2)

	if idx != -1 {
		fmt.Printf("Found at index %d in %d steps\n", idx, steps)
	} else {
		fmt.Printf("Not found after %d steps\n", steps)
	}

}

func linearSteps(array []int, searchValue int) (int, int) {
	steps := 0
	for index, element := range array {
		steps++ // count every comparison

		if element == searchValue {
			return index, steps
		}
		if element > searchValue {
			break
		}
	}
	return -1, steps
}
