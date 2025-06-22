package main

func isLeapYear(year int) bool {
	if year%100 == 0 {
		year%400 == 0
	}
	return year%4 == 0
}

// Time Complexity: O(1)
// No matter size of the input, it always takes Two Steps
