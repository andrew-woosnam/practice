package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{arr: []int{1, 2, 3, 4, 5}, target: 3, expected: 2},
		{arr: []int{1, 2, 3, 4, 5}, target: 6, expected: -1},
		{arr: []int{1, 2, 3, 4, 5}, target: 1, expected: 0},
		{arr: []int{1, 2, 3, 4, 5}, target: 5, expected: 4},
		{arr: []int{}, target: 1, expected: -1},
		{arr: []int{1}, target: 1, expected: 0},
		{arr: []int{1}, target: 0, expected: -1},
	}

	for i, test := range tests {
		result := BinarySearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("Test %d failed: got %d, expected %d", i+1, result, test.expected)
		}
	}
}
