package reverseinteger

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Positive number", 123, 321},
		{"Negative number", -123, -321},
		{"Number with trailing zeros", 120, 21},
		{"Single digit number", 5, 5},
		{"Zero", 0, 0},
		{"Overflow positive", 1534236469, 0},
		{"Overflow negative", -1534236469, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.input)
			if result != tt.expected {
				t.Errorf("reverse(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(123456789)
	}
}
