package main

import (
	"fmt"
)

func even_numbers_v1() {
	number := 2
	for number <= 100 {
		// If number is even, print it:
		if number%2 == 0 {
			fmt.Println(number)
		}
		number++
	}
}

func even_numbers_v2() {
	number := 2
	for number <= 100 {
		fmt.Println(number)
		number += 2
	}
}
