package main

import "fmt"

func mergeF(arr1, arr2 []int) []int {
	var newArray []int
	arr1_point := 0
	arr2_point := 0

	for {
		switch {
		// case 1: both arrays are exhausted
		case arr1_point >= len(arr1) && arr2_point >= len(arr2):
			return newArray

		case arr1_point >= len(arr1):
			newArray = append(newArray, arr2[arr2_point])
			arr2_point++

		case arr2_point >= len(arr2):
			newArray = append(newArray, arr1[arr1_point])
			arr1_point++

		case arr1[arr1_point] < arr2[arr2_point]:
			newArray = append(newArray, arr1[arr1_point])
			arr1_point++

		default:
			newArray = append(newArray, arr2[arr2_point])
			arr2_point++
		}
	}
}

func merge(arr1, arr2 []int) []int {
	var newArray []int
	arr1_point := 0
	arr2_point := 0

	for arr1_point < len(arr1) && arr2_point < len(arr2) {
		if arr1[arr1_point] < arr2[arr2_point] {
			newArray = append(newArray, arr1[arr1_point])
			arr1_point++
		} else {
			newArray = append(newArray, arr2[arr2_point])
			arr2_point++
		}
	}

	for arr1_point < len(arr1) {
		newArray = append(newArray, arr1[arr1_point])
		arr1_point++
	}

	for arr2_point < len(arr2) {
		newArray = append(newArray, arr2[arr2_point])
		arr2_point++

	}
	return newArray
}

func main() {
	array1 := []int{1, 3, 5, 7}
	array2 := []int{2, 4, 6}
	finalArray := merge(array1, array2)
	finalFArray := mergeF(array1, array2)

	fmt.Println("New Array:", finalArray)
	fmt.Println("New Array with Case st:", finalFArray)
}
