package main

type sortableArray []int

func (arr sortableArray) partition(leftPtr, rightPtr int) int {
	// We always choose the right-most element as the pivot.
	// We keep the index of the pivot for later use.
	pivotIdx := rightPtr

	// We grab the pivot value itself
	pivot := arr[pivotIdx]

	// We start the right pointer immediately to the left of the pivot
	rightPtr = rightPtr - 1

	for {
		// Move the left pointer to the right as long as it points to value
		// that is less than the pivot.
		for arr[leftPtr] < pivot {
			leftPtr = leftPtr + 1
		}

		// Move the right pointer to the left as long as it points to a
		// value that is greater than the pivot.
		for arr[rightPtr] > pivot {
			rightPtr = rightPtr - 1
		}

		// We've now reached the point where we've stopped moving both the
		// left and right pointers.
		//
		// We check whether the left pointer has reached or gone beyond the
		// right pointer. If it has, we break out of the loop so we can
		// swap the pivot later in our code.
		if leftPtr >= rightPtr {
			break
		} else {
			// If the left pointer is still to the left of the right
			// pointer, we swap the values of the left and right pointers.
			arr[leftPtr], arr[rightPtr] = arr[rightPtr], arr[leftPtr]

			// We move the left pointer over to the right, gearing up for
			// the next round of left and right pointer movements.
			leftPtr = leftPtr + 1
		}
	}
	// As the final step of the partition, we swap the value of the
	// left pointer with the pivot.
	arr[leftPtr], arr[pivotIdx] = arr[pivotIdx], arr[leftPtr]

	return leftPtr
}
