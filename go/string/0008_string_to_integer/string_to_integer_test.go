package stringtointeger

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected int
	}{
		{"negative", "-45", -45},
		{"positive", "+32", 32},
		{"space", "   64", 64},

		{"words after", "4193 with words", 4193},
		{"words before", "words and 987", 0},
		{"multiple signs", "+-12", 0},
		{"only sign", "-", 0},
		{"empty", "", 0},
		{"only spaces", "     ", 0},

		{"leading zeros", "00000000000123", 123},
		{"overflow positive", "2147483648", math.MaxInt32},
		{"overflow negative", "-91283472332", math.MinInt32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := myAtoi(tt.value); result != tt.expected {
				t.Errorf("myAtoi(%q) = %d; esperado %d", tt.value, result, tt.expected)
			}
		})
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	input := "   -4193 with words"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		myAtoi(input)
	}
}
