package main

func chessBoardSpace(int numberOfGrains) int {
	chessBoardSpaces := 1
	placedGrains := 1

	for placedGrains < numberOfGrains {
		placedGrains *= 2
		chessBoardSpaces += 1
	}
	return chessBoardSpaces
}

// O(log N): For each chessBoardSpaces, the number of grains will
// become halve of previous i.e. decrement of halve of the previous
// input
