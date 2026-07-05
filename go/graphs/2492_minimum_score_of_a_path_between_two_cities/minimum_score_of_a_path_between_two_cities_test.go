package minimumscoreofapathbetweentwocities

import (
	"testing"
)

func TestMinScore(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		roads    [][]int
		expected int
	}{
		{
			name:     "Minimum edge 5",
			n:        4,
			roads:    [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			expected: 5,
		},
		{
			name:     "Minimum edge 2 at start",
			n:        4,
			roads:    [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}},
			expected: 2,
		},
		{
			name:     "Dead end with the minimum road",
			n:        4,
			roads:    [][]int{{1, 4, 100}, {1, 2, 50}, {2, 3, 3}},
			expected: 3,
		},
		{
			name:     "Isolated component should be ignored",
			n:        5,
			roads:    [][]int{{1, 2, 10}, {2, 4, 8}, {3, 5, 1}},
			expected: 8,
		},
		{
			name:     "Simple linear path",
			n:        3,
			roads:    [][]int{{1, 2, 100}, {2, 3, 10}},
			expected: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := minScore(tc.n, tc.roads)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkMinScore(b *testing.B) {
	n := 1000
	roads := make([][]int, 0, n-1)

	for i := 1; i < n; i++ {
		distance := 10000 - i
		roads = append(roads, []int{i, i + 1, distance})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		minScore(n, roads)
	}
}
