package main

import "fmt"

// InsertionSort
func InsertionSort(array []int) {
	for index := 1; index < len(array); index++ {
		key := array[index]
		position := index - 1

		for position >= 0 && array[position] > key {
			array[position+1] = array[position]
			position = position - 1
		}
		array[position+1] = key
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Original array:", arr)

	InsertionSort(arr)

	fmt.Println("Sorted array:", arr)
}
