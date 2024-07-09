package binarysearch

import "testing"

func TestBinarySearchIterative_EmptyArray(t *testing.T) {
	result := BinarySearchIterative([]int{}, 1)
	if result != -1 {
		t.Errorf("Expected -1 for empty array, got %d", result)
	}
}

func TestBinarySearchIterative_SingleElementFound(t *testing.T) {
	result := BinarySearchIterative([]int{1}, 1)
	if result != 0 {
		t.Errorf("Expected 0 for single element array where element is found, got %d", result)
	}
}

func TestBinarySearchIterative_SingleElementNotFound(t *testing.T) {
	result := BinarySearchIterative([]int{1}, 0)
	if result != -1 {
		t.Errorf("Expected -1 for single element array where element is not found, got %d", result)
	}
}

func TestBinarySearchIterative_TwoElementsFirstElementFound(t *testing.T) {
	result := BinarySearchIterative([]int{1, 2}, 1)
	if result != 0 {
		t.Errorf("Expected 0 for two element array where first element is found, got %d", result)
	}
}

func TestBinarySearchIterative_TwoElementsSecondElementFound(t *testing.T) {
	result := BinarySearchIterative([]int{1, 2}, 2)
	if result != 1 {
		t.Errorf("Expected 1 for two element array where second element is found, got %d", result)
	}
}

func TestBinarySearchIterative_TwoElementsNotFound(t *testing.T) {
	result := BinarySearchIterative([]int{1, 2}, 3)
	if result != -1 {
		t.Errorf("Expected -1 for two element array where element is not found, got %d", result)
	}
}

func TestBinarySearchIterative_MultipleElementsFound(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{arr: []int{1, 2, 3, 4, 5}, target: 3, expected: 2},
		{arr: []int{1, 2, 3, 4, 5}, target: 1, expected: 0},
		{arr: []int{1, 2, 3, 4, 5}, target: 5, expected: 4},
	}

	for _, test := range tests {
		result := BinarySearchIterative(test.arr, test.target)
		if result != test.expected {
			t.Errorf("For array %v and target %d, expected %d, got %d", test.arr, test.target, test.expected, result)
		}
	}
}

func TestBinarySearchRecursive_EmptyArray(t *testing.T) {
	result := BinarySearchRecursive([]int{}, 1)
	if result != -1 {
		t.Errorf("Expected -1 for empty array, got %d", result)
	}
}

func TestBinarySearchRecursive_SingleElementFound(t *testing.T) {
	result := BinarySearchRecursive([]int{1}, 1)
	if result != 0 {
		t.Errorf("Expected 0 for single element array where element is found, got %d", result)
	}
}

func TestBinarySearchRecursive_SingleElementNotFound(t *testing.T) {
	result := BinarySearchRecursive([]int{1}, 0)
	if result != -1 {
		t.Errorf("Expected -1 for single element array where element is not found, got %d", result)
	}
}

func TestBinarySearchRecursive_TwoElementsFirstElementFound(t *testing.T) {
	result := BinarySearchRecursive([]int{1, 2}, 1)
	if result != 0 {
		t.Errorf("Expected 0 for two element array where first element is found, got %d", result)
	}
}

func TestBinarySearchRecursive_TwoElementsSecondElementFound(t *testing.T) {
	result := BinarySearchRecursive([]int{1, 2}, 2)
	if result != 1 {
		t.Errorf("Expected 1 for two element array where second element is found, got %d", result)
	}
}

func TestBinarySearchRecursive_TwoElementsNotFound(t *testing.T) {
	result := BinarySearchRecursive([]int{1, 2}, 3)
	if result != -1 {
		t.Errorf("Expected -1 for two element array where element is not found, got %d", result)
	}
}

func TestBinarySearchRecursive_MultipleElementsFound(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{arr: []int{1, 2, 3, 4, 5}, target: 3, expected: 2},
		{arr: []int{1, 2, 3, 4, 5}, target: 1, expected: 0},
		{arr: []int{1, 2, 3, 4, 5}, target: 5, expected: 4},
	}

	for _, test := range tests {
		result := BinarySearchRecursive(test.arr, test.target)
		if result != test.expected {
			t.Errorf("For array %v and target %d, expected %d, got %d", test.arr, test.target, test.expected, result)
		}
	}
}
