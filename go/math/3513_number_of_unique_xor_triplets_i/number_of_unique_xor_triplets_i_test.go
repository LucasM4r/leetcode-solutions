package numberofuniquexortripletsi

import (
	"testing"
)

func TestUniqueXorTriplets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"EmptySlice", []int{}, 0},
		{"SingleElement", []int{1}, 1},
		{"TwoElements", []int{1, 2}, 2},
		{"ThreeElements", []int{1, 2, 3}, 4},
		{"FourElements", []int{1, 2, 3, 4}, 8},
		{"FiveElements", []int{1, 2, 3, 4, 5}, 8},
		{"SevenElements", []int{1, 2, 3, 4, 5, 6, 7}, 8},
		{"EightElements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 16},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := uniqueXorTriplets(tc.nums)
			if result != tc.expected {
				t.Errorf("uniqueXorTriplets(%v) = %v; want %v", tc.nums, result, tc.expected)
			}
		})
	}
}

func BenchmarkUniqueXorTripletsLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uniqueXorTriplets(nums)
	}
}
