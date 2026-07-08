package concatenatenonzerodigitsandmultiplybysumi

import "testing"

func TestSumAndMultiply(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int64
	}{
		{"Basic case 123", 123, 738},
		{"Case with zero 102", 102, 36},
		{"Case with multiple zeros 505", 505, 550},
		{"Zero input", 0, 0},
		{"All zeros 100", 100, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumAndMultiply(tt.input)
			if result != tt.expected {
				t.Errorf("sumAndMultiply(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkSumAndMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumAndMultiply(120340066400056789)
	}
}
