package chapter_01

import (
	"fmt"
)

func main() {
	// print_numbers_version_one()
	print_numbers_version_two()
}

func print_numbers_version_one() {
	number := 2
	for number <= 100 {
		// If number is even, print it:
		if number%2 == 0 {
			fmt.Println(number)
		}
		number++
	}
}

func print_numbers_version_two() {
	number := 2
	for number <= 100 {
		fmt.Println(number)
		number += 2
	}
}
