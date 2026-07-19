package findgreatestcommondivisorofarray

import (
	"testing"
)

func TestFindGCD(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Standard array with GCD > 1",
			nums:     []int{2, 5, 6, 9, 10},
			expected: 2,
		},
		{
			name:     "Standard array with GCD = 1",
			nums:     []int{7, 5, 6, 8, 3},
			expected: 1,
		},
		{
			name:     "Array with identical elements",
			nums:     []int{3, 3},
			expected: 3,
		},
		{
			name:     "Extremes of constraints",
			nums:     []int{1, 1000},
			expected: 1,
		},
		{
			name:     "Multiple minimums and maximums",
			nums:     []int{12, 12, 18, 15, 24, 24},
			expected: 12,
		},
		{
			name:     "Prime numbers",
			nums:     []int{17, 3, 5, 13, 23},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.nums); got != tt.expected {
				t.Errorf("findGCD(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}

func BenchmarkFindGCD(b *testing.B) {
	nums := []int{
		2, 5, 6, 9, 10, 15, 20, 25, 30, 35,
		40, 45, 50, 55, 60, 65, 70, 75, 80, 85,
		90, 95, 100, 1, 1000, 42, 68, 99, 101, 777,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findGCD(nums)
	}
}
