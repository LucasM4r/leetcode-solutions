package longestcommonprefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "common prefix exists",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "no common prefix",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "empty slice",
			strs:     []string{},
			expected: "",
		},
		{
			name:     "single string in slice",
			strs:     []string{"hello"},
			expected: "hello",
		},
		{
			name:     "all strings identical",
			strs:     []string{"test", "test", "test"},
			expected: "test",
		},
		{
			name:     "contains empty string",
			strs:     []string{"ab", "", "c"},
			expected: "",
		},
		{
			name:     "shortest string is the prefix",
			strs:     []string{"a", "ab", "abc"},
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %v, expected %v", tt.strs, got, tt.expected)
			}
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	benchmarks := []struct {
		name string
		strs []string
	}{
		{
			name: "short strings match",
			strs: []string{"flower", "flow", "flight"},
		},
		{
			name: "no match",
			strs: []string{"dog", "racecar", "car"},
		},
		{
			name: "identical strings",
			strs: []string{"test", "test", "test", "test", "test"},
		},
		{
			name: "empty slice",
			strs: []string{},
		},
		{
			name: "longer strings",
			strs: []string{
				"understanding_benchmark_testing_in_go",
				"understanding_benchmark_functions",
				"understanding_benchmarks",
			},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				longestCommonPrefix(bm.strs)
			}
		})
	}
}
