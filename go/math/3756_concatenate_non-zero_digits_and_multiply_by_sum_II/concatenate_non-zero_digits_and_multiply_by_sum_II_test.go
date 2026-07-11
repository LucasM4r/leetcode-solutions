package concatenatenonzerodigitsandmultiplybysumii

import (
	"reflect"
	"strings"
	"testing"
)

func TestSumAndMultiply(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		queries  [][]int
		expected []int
	}{
		{
			name:     "Standard Scenario",
			s:        "10203",
			queries:  [][]int{{0, 4}, {0, 1}, {2, 4}},
			expected: []int{738, 1, 115},
		},
		{
			name:     "Only Zeros In String",
			s:        "00000",
			queries:  [][]int{{0, 4}, {1, 3}},
			expected: []int{0, 0},
		},
		{
			name:     "Empty String",
			s:        "",
			queries:  [][]int{{0, 0}},
			expected: []int{},
		},
		{
			name:     "Empty Queries",
			s:        "123",
			queries:  [][]int{},
			expected: []int{},
		},
		{
			name:     "No Non-Zero Digits In Range",
			s:        "10002",
			queries:  [][]int{{1, 3}},
			expected: []int{0},
		},
		{
			name:     "Long String Without Zeros",
			s:        "12345",
			queries:  [][]int{{0, 2}},
			expected: []int{738},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumAndMultiply(tt.s, tt.queries)

			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkSumAndMultiply(b *testing.B) {
	s := strings.Repeat("1020304050", 1000)

	var queries [][]int
	for i := 0; i < 1000; i++ {
		start := i * 5
		end := start + 500
		if end >= len(s) {
			end = len(s) - 1
		}
		queries = append(queries, []int{start, end})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sumAndMultiply(s, queries)
	}
}
