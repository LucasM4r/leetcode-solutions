package sumofgcdofformedpairs

import (
	"math/rand"
	"testing"
)

func TestGcdSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Example 1 (Odd length)",
			nums:     []int{2, 6, 4},
			expected: 2,
		},
		{
			name:     "Example 2 (Even length)",
			nums:     []int{3, 6, 2, 8},
			expected: 5,
		},
		{
			name:     "Single element (Should return 0 as no pairs can be formed)",
			nums:     []int{10},
			expected: 0,
		},
		{
			name:     "Two identical elements",
			nums:     []int{5, 5},
			expected: 5,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{12, 8, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			actual := gcdSum(numsCopy)
			if actual != tt.expected {
				t.Errorf("gcdSum(%v) = %d; want %d", tt.nums, actual, tt.expected)
			}
		})
	}
}

func helperGenerateRandomSlice(size int) []int {
	rng := rand.New(rand.NewSource(42))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rng.Intn(1_000_000_000) + 1
	}
	return slice
}

func BenchmarkGcdSum(b *testing.B) {
	nums := helperGenerateRandomSlice(100_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		_ = gcdSum(input)
	}
}
