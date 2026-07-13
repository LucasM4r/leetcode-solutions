package ranktransformofanarray

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "EmptyArray",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "DistinctElements",
			input:    []int{40, 10, 20, 30},
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "AllIdenticalElements",
			input:    []int{100, 100, 100},
			expected: []int{1, 1, 1},
		},
		{
			name:     "MixedElementsWithDuplicates",
			input:    []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			expected: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			name:     "AlreadySorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "ReverseSorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "WithNegativeNumbers",
			input:    []int{-10, 0, 10, -20},
			expected: []int{2, 3, 4, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			result := arrayRankTransform(inputCopy)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ArrayRankTransform() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func benchmarkHelper(size int, b *testing.B) {
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = i % (size/2 + 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrayRankTransform(input)
	}
}

func BenchmarkArrayRankTransform(b *testing.B) {
	benchmarkHelper(100000, b)
}
