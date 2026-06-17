package generateparenthesis

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "One pair n=1",
			n:        1,
			expected: []string{"()"},
		},
		{
			name:     "Two pairs n=2",
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			name:     "Three pairs n=3",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "Zero pairs n=0",
			n:        0,
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.n); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("generateParenthesis(%d) = %v; want %v", tt.n, got, tt.expected)
			}
		})
	}
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(8)
	}
}
