package main

import "fmt"

func duplicate(arr []string) string {
	m := make(map[string]bool)
	var dupe string

	for _, i := range arr {
		if !m[i] {
			m[i] = true
		} else {
			dupe = i
		}
	}
	return dupe
}

func main() {
	array := []string{"a", "b", "c", "d", "c", "e", "f"}
	res := duplicate(array)
	fmt.Println(res)
}
