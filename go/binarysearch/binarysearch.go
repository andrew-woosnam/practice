package binarysearch

// performs binary search on a sorted array, returning the target's index if found, else -1
func BinarySearch(collection []int, target int) int {
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
