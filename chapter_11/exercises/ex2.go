package main

import "fmt"

func getEvens(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	evens := getEvens(nums[1:])

	if nums[0]%2 == 0 {
		return append([]int{nums[0]}, evens...)
	}
	return evens
}

// (define (get-evens lon)
//   (cond
//     [(empty? lon) '()]
//     [else (if (even? (first lon))
//               (cons (first lon) (get-evens (rest lon)))
//               (get-evens (rest lon)))]))

func main() {
	fmt.Println(getEvens([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
