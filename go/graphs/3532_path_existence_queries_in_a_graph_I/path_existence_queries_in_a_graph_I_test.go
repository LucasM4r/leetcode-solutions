package pathexistencequeriesinagraphi

import (
	"reflect"
	"testing"
)

func TestPathExistenceQueries(t *testing.T) {
	tests := []struct {
		n        int
		nums     []int
		maxDiff  int
		queries  [][]int
		expected []bool
	}{
		{
			n:        4,
			nums:     []int{1, 3, 8, 9},
			maxDiff:  3,
			queries:  [][]int{{0, 1}, {0, 2}, {2, 3}},
			expected: []bool{true, false, true},
		},
		{
			n:        5,
			nums:     []int{1, 2, 3, 4, 5},
			maxDiff:  1,
			queries:  [][]int{{0, 4}, {1, 2}},
			expected: []bool{true, true},
		},
		{
			n:        3,
			nums:     []int{1, 10, 20},
			maxDiff:  5,
			queries:  [][]int{{0, 1}, {1, 2}, {0, 2}},
			expected: []bool{false, false, false},
		},
	}

	for _, tt := range tests {
		actual := pathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("pathExistenceQueries(%d, %v, %d, %v) = %v; want %v", tt.n, tt.nums, tt.maxDiff, tt.queries, actual, tt.expected)
		}
	}
}

func BenchmarkPathExistenceQueries(b *testing.B) {
	n := 1000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i * 2
	}
	queries := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		queries[i] = []int{i % n, (i + 5) % n}
	}
	maxDiff := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pathExistenceQueries(n, nums, maxDiff, queries)
	}
}
