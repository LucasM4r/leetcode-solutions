package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		k        int
	}{
		{"Example 1", []int{1, 1, 2}, []int{1, 2}, 2},
		{"Example 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
		{"Empty slice", []int{}, []int{}, 0},
		{"No duplicates", []int{1, 2, 3}, []int{1, 2, 3}, 3},
		{"All same", []int{1, 1, 1, 1}, []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)

			k := removeDuplicates(nums)

			if k != tt.k {
				t.Errorf("expected k=%d, got %d", tt.k, k)
			}
			if !reflect.DeepEqual(nums[:k], tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, nums[:k])
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, len(input))
		copy(data, input)
		removeDuplicates(data)
	}
}
