package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Single valid pair", "()", true},
		{"Multiple valid pairs", "()[]{}", true},
		{"Mismatched pair", "(]", false},
		{"Interleaved invalid", "([)]", false},
		{"Nested valid", "{[]}", true},
		{"Only opening", "((((", false},
		{"Only closing", "))))", false},
		{"Empty string", "", true},
		{"Long nested valid", "({[()()]})", true},
		{"Long invalid", "({[()()]}}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	input := "({[()()]})({[()()]})({[()()]})({[()()]})"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValid(input)
	}
}
