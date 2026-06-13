package findthetownjudge

import "testing"

func TestFindJudge(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		trust    [][]int
		expected int
	}{
		{
			name:     "Basic two person case",
			n:        2,
			trust:    [][]int{{1, 2}},
			expected: 2,
		},
		{
			name:     "Judge with multiple followers",
			n:        3,
			trust:    [][]int{{1, 3}, {2, 3}},
			expected: 3,
		},
		{
			name:     "No judge due to cycle",
			n:        3,
			trust:    [][]int{{1, 3}, {2, 3}, {3, 1}},
			expected: -1,
		},
		{
			name:     "Complex relationship network",
			n:        4,
			trust:    [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}},
			expected: 3,
		},
		{
			name:     "Single person as judge",
			n:        1,
			trust:    [][]int{},
			expected: 1,
		},
		{
			name:     "No trust relationships",
			n:        2,
			trust:    [][]int{},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findJudge(tt.n, tt.trust)
			if result != tt.expected {
				t.Errorf("FindJudge() = %d; want %d", result, tt.expected)
			}
		})
	}
}

func BenchmarkFindJudge(b *testing.B) {
	n := 1000
	trust := [][]int{{1, 2}, {3, 4}, {5, 6}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findJudge(n, trust)
	}
}
