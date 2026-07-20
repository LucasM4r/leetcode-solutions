package smallestsubsequenceofdistinctcharacters

import "testing"

func TestSmallestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "Example 2",
			s:    "cbacdcbc",
			want: "acdb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.s); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSmallestSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestSubsequence("cbacdcbc")
	}
}
