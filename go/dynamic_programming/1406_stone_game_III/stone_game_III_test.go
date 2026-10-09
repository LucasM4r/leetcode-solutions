package stonegameiii

import (
	"testing"
)

func TestStoneGameIII(t *testing.T) {
	tests := []struct {
		name       string
		stoneValue []int
		expected   string
	}{
		{
			name:       "Bob wins",
			stoneValue: []int{1, 2, 3, 7},
			expected:   "Bob",
		},
		{
			name:       "Alice wins with negative values at the end",
			stoneValue: []int{1, 2, 3, -9},
			expected:   "Alice",
		},
		{
			name:       "Tie game",
			stoneValue: []int{1, 2, 3, 6},
			expected:   "Tie",
		},
		{
			name:       "Alice wins with single stone",
			stoneValue: []int{10},
			expected:   "Alice",
		},
		{
			name:       "All negative stones",
			stoneValue: []int{-1, -2, -3},
			expected:   "Tie",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stoneGameIII(tt.stoneValue)
			if result != tt.expected {
				t.Errorf("stoneGameIII(%v) = %q; want %q", tt.stoneValue, result, tt.expected)
			}
		})
	}
}

func BenchmarkStoneGameIII(b *testing.B) {
	input := []int{1, 2, 3, 7, 1, 2, 3, -9, 10, -1, -2, -3, 5, 5, 5, 10, 15, 20}

	for b.Loop() {
		stoneGameIII(input)
	}
}
