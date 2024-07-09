package binarysearch

import (
	"context"
	"testing"
	"time"
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
		t.Run("", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			done := make(chan struct{})
			go func() {
				result := BinarySearch(test.arr, test.target)
				if result != test.expected {
					t.Errorf("Test %d failed: got %d, expected %d", i+1, result, test.expected)
				}
				done <- struct{}{}
			}()

			select {
			case <-ctx.Done():
				t.Errorf("Test %d timed out", i+1)
			case <-done:
			}
		})
	}
}
