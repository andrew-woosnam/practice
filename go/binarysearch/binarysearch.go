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
	return recurse(collection, target, 0, len(collection)-1)
}

func recurse(collection []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	middle := start + (end-start)/2 // find mid without risking int overflow
	sample := collection[middle]
	if sample == target {
		return middle
	} else if sample < target {
		return recurse(collection, target, middle+1, end)
	} else { // sample > target
		return recurse(collection, target, start, middle-1)
	}
}
