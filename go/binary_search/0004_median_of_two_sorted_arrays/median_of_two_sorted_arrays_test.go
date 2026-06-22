package medianoftwosortedarrays

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "Example 1 odd total",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			name:     "Example 2 even total",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			name:     "One empty array",
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			name:     "Different sizes with negatives",
			nums1:    []int{-5, -3, -1},
			nums2:    []int{2, 4, 6, 8, 10},
			expected: 3.0,
		},
		{
			name:     "All zeros",
			nums1:    []int{0, 0},
			nums2:    []int{0, 0},
			expected: 0.0,
		},
		{
			name:     "Single values each array",
			nums1:    []int{1},
			nums2:    []int{2},
			expected: 1.5,
		},
		{
			name:     "Unbalanced sizes",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{100},
			expected: 3.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v; expected %v", tt.nums1, tt.nums2, got, tt.expected)
			}
		})
	}
}

func BenchmarkFindMedianSortedArraysBalanced(b *testing.B) {
	nums1 := makeSortedEvenNumbers(100000)
	nums2 := makeSortedOddNumbers(100000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
}

func BenchmarkFindMedianSortedArraysSkewed(b *testing.B) {
	nums1 := makeSortedEvenNumbers(100)
	nums2 := makeSortedOddNumbers(1000000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
}

func makeSortedEvenNumbers(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i * 2
	}
	return arr
}

func makeSortedOddNumbers(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i*2 + 1
	}
	return arr
}
