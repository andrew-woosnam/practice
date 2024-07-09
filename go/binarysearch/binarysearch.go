package binarysearch

// performs iterative binary search on a sorted array, returning the target's index if found, else -1
func BinarySearchIterative(collection []int, target int) int {
	l := 0
	r := len(collection) - 1

	for l <= r {
		middle := (l + r) / 2
		sample := collection[middle]

		if sample == target {
			return middle
		}

		if target > sample {
			l = middle + 1
		} else if target < sample {
			r = middle - 1
		}
	}

	return -1
}

// performs recursive binary search on a sorted array, returning the target's index if found, else -1
func BinarySearchRecursive(collection []int, target int) int {
	if len(collection) == 0 {
		return -1
	}

	middle := len(collection) / 2
	sample := collection[middle]
	if sample == target {
		return middle
	}

	if target > sample {
		newStart := middle + 1
		result := BinarySearchRecursive(collection[newStart:], target)
		if result == -1 {
			return -1
		}
		return newStart + result
	}

	// target < sample
	return BinarySearchRecursive(collection[:middle], target)
}
