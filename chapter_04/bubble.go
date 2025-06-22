package main

import "fmt"

func bubbleSort(list []int) []int {
	unsortedIndex := len(list) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsortedIndex; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}
		}
		unsortedIndex -= 1
	}
	return list
}

func main() {
	numbers := []int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println(bubbleSort(numbers))
}
