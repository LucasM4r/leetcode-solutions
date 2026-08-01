package findthesafestpathinagrid

import (
	"testing"
)

func TestMaximumSafenessFactor(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Thieves at start and end",
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 0,
		},
		{
			name: "Thief at top right corner",
			grid: [][]int{
				{0, 0, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "Larger grid with multiple thieves",
			grid: [][]int{
				{0, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "All thieves",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 0,
		},
		{
			name: "Path strictly blocked by thieves",
			grid: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 0},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumSafenessFactor(tt.grid)
			if result != tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}

func BenchmarkMaximumSafenessFactor(b *testing.B) {
	grid := [][]int{
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maximumSafenessFactor(grid)
	}
}
