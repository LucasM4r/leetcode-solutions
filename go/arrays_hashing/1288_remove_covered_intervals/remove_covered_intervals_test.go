package removecoveredintervals

import (
	"testing"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Example case 1",
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			expected:  2,
		},
		{
			name:      "Example case 2",
			intervals: [][]int{{1, 4}, {2, 3}},
			expected:  1,
		},
		{
			name:      "Empty intervals",
			intervals: [][]int{},
			expected:  0,
		},
		{
			name:      "Single interval",
			intervals: [][]int{{1, 5}},
			expected:  1,
		},
		{
			name:      "Identical intervals",
			intervals: [][]int{{1, 4}, {1, 4}, {1, 4}},
			expected:  1,
		},
		{
			name:      "Same start, different end",
			intervals: [][]int{{1, 2}, {1, 4}, {1, 3}},
			expected:  1,
		},
		{
			name:      "No overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected:  3,
		},
		{
			name:      "Disjoint but contiguous",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			expected:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intervalsCopy := make([][]int, len(tt.intervals))
			for i, v := range tt.intervals {
				intervalsCopy[i] = append([]int(nil), v...)
			}

			actual := removeCoveredIntervals(intervalsCopy)
			if actual != tt.expected {
				t.Errorf("removeCoveredIntervals() = %d; want %d", actual, tt.expected)
			}
		})
	}
}

func BenchmarkRemoveCoveredIntervals_Small(b *testing.B) {
	intervals := [][]int{{1, 4}, {3, 6}, {2, 8}, {1, 5}, {4, 9}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([][]int, len(intervals))
		for j, v := range intervals {
			input[j] = append([]int(nil), v...)
		}

		_ = removeCoveredIntervals(input)
	}
}

func BenchmarkRemoveCoveredIntervals_Large(b *testing.B) {
	size := 1000
	intervals := make([][]int, size)
	for i := 0; i < size; i++ {
		intervals[i] = []int{i, i + 10}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([][]int, len(intervals))
		for j, v := range intervals {
			input[j] = append([]int(nil), v...)
		}

		_ = removeCoveredIntervals(input)
	}
}
