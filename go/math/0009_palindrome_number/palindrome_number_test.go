package palindromenumber

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"Example 1 (Odd length)", 121, true},
		{"Example 2 (Negative)", -121, false},
		{"Example 3 (Ends with 0)", 10, false},
		{"Single digit zero", 0, true},
		{"Single digit non-zero", 7, true},
		{"Even length palindrome", 1221, true},
		{"Long odd length palindrome", 1234321, true},
		{"Not a palindrome", 123, false},
		{"Starts and ends with same but not palindrome", 1000021, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.value); got != tt.expected {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.value, got, tt.expected)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	benchmarks := []struct {
		name  string
		value int
	}{
		{"SingleDigit", 7},
		{"Negative", -121},
		{"EndsWithZero", 1000},
		{"EvenLength", 1221},
		{"OddLength_Short", 121},
		{"OddLength_Long", 1234321},
		{"NotPalindrome_Long", 1234567},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPalindrome(bm.value)
			}
		})
	}
}
