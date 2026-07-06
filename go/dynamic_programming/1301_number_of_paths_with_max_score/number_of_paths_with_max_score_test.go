package numberofpathswithmaxscore

import (
	"reflect"
	"testing"
)

func TestPathsWithMaxScore(t *testing.T) {
	tests := []struct {
		name     string
		board    []string
		expected []int
	}{
		{
			name:     "Single Path",
			board:    []string{"E23", "2X2", "12S"},
			expected: []int{7, 1},
		},
		{
			name:     "Multiple Paths",
			board:    []string{"E12", "1X1", "21S"},
			expected: []int{4, 2},
		},
		{
			name:     "No Possible Path (Blocked)",
			board:    []string{"E11", "XXX", "11S"},
			expected: []int{0, 0},
		},
		{
			name:     "Larger Grid without Blocks",
			board:    []string{"E99", "999", "99S"},
			expected: []int{27, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pathsWithMaxScore(tt.board)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkPathsWithMaxScore(b *testing.B) {
	board := []string{
		"E999999999",
		"9999999999",
		"9999999999",
		"999X999X99",
		"9999999999",
		"9999999999",
		"99X9999999",
		"9999999999",
		"9999999999",
		"999999999S",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pathsWithMaxScore(board)
	}
}
