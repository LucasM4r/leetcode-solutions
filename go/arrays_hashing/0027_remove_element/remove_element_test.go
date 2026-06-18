package removeelement

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{
			name:     "Empty slice",
			nums:     []int{},
			val:      3,
			expected: 0,
		},
		{
			name:     "Value not present",
			nums:     []int{1, 2, 4, 5},
			val:      3,
			expected: 4,
		},
		{
			name:     "Value present once",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			name:     "Value present multiple times",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
		{
			name:     "All elements match value",
			nums:     []int{2, 2, 2},
			val:      2,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			if got := removeElement(numsCopy, tt.val); got != tt.expected {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.nums, tt.val, got, tt.expected)
			}
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 5, 2, 7, 8, 2, 9, 10}
	val := 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		temp := make([]int, len(nums))
		copy(temp, nums)

		removeElement(temp, val)
	}
}
