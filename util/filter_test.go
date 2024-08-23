package util

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	result := Filter(numbers, func(a int) bool { return a%2 == 0 })
	if slices.Contains(result, 1) || slices.Contains(result, 3) {
		t.Fatalf(`Filtering out odd numbers = %v, %v, want %v`, result, "Error filtering slice", []int{2, 4})
	}

}
