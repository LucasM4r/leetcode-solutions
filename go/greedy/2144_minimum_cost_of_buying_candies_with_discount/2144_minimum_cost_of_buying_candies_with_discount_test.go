package minimumcostofbuyingcandieswithdiscount

import (
	"testing"
)

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			cost:     []int{1, 2, 3},
			expected: 5,
		},
		{
			name:     "Example 2",
			cost:     []int{6, 5, 7, 9, 2, 2},
			expected: 23,
		},
		{
			name:     "Only 2 items",
			cost:     []int{5, 5},
			expected: 10,
		},
		{
			name:     "Empty list",
			cost:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost(tt.cost)
			if got != tt.expected {
				t.Errorf("MinimumCost() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func BenchmarkMinimumCost(b *testing.B) {
	data := []int{6, 5, 7, 9, 2, 2, 10, 20, 30, 40, 50, 60}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		temp := make([]int, len(data))
		copy(temp, data)
		minimumCost(temp)
	}
}
