package sequentialdigits

import (
	"reflect"
	"testing"
)

func TestSequentialDigits(t *testing.T) {
	tests := []struct {
		name     string
		low      int
		high     int
		expected []int
	}{
		{
			name:     "Example 1",
			low:      100,
			high:     300,
			expected: []int{123, 234},
		},
		{
			name:     "Example 2",
			low:      1000,
			high:     13000,
			expected: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
		{
			name:     "No sequential digits in range",
			low:      50,
			high:     55,
			expected: nil,
		},
		{
			name:     "Only single digit numbers (function focuses on length >= 2)",
			low:      1,
			high:     9,
			expected: nil,
		},
		{
			name: "Maximum range",
			low:  10,
			high: 1000000000,
			expected: []int{
				12, 23, 34, 45, 56, 67, 78, 89,
				123, 234, 345, 456, 567, 678, 789,
				1234, 2345, 3456, 4567, 5678, 6789,
				12345, 23456, 34567, 45678, 56789,
				123456, 234567, 345678, 456789,
				1234567, 2345678, 3456789,
				12345678, 23456789,
				123456789,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sequentialDigits(tt.low, tt.high)

			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sequentialDigits(%d, %d) = %v; expected %v", tt.low, tt.high, result, tt.expected)
			}
		})
	}
}

func BenchmarkSequentialDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sequentialDigits(1000, 13000)
	}
}

func BenchmarkSequentialDigitsMaxRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sequentialDigits(10, 1000000000)
	}
}
