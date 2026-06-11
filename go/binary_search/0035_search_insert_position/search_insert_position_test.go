package searchinsertposition

import (
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		target   int
		expected int
	}{
		{
			name:     "Target exists in the middle",
			array:    []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Target does not exist, insert in the middle",
			array:    []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Target does not exist, insert at the end",
			array:    []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "Target does not exist, insert at the beginning",
			array:    []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "Single element array, target is smaller",
			array:    []int{1},
			target:   0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.array, tt.target)

			if got != tt.expected {
				t.Errorf("searchInsert(%v, %d) = %d; expected %d", tt.array, tt.target, got, tt.expected)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	size := 100000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i * 2
	}
	target := size * 2

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searchInsert(arr, target)
	}
}
