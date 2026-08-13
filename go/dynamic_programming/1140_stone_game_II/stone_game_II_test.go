package stonegameii

import (
	"testing"
)

func TestStoneGameII(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  int
	}{
		{
			name:  "Example 1",
			piles: []int{2, 7, 9, 4, 4},
			want:  10,
		},
		{
			name:  "Example 2",
			piles: []int{1, 2, 3, 4, 5, 100},
			want:  104,
		},
		{
			name:  "Single pile",
			piles: []int{5},
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameII(tt.piles); got != tt.want {
				t.Errorf("stoneGameII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStoneGameII(b *testing.B) {
	piles := []int{2, 7, 9, 4, 4, 1, 5, 6, 8, 10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stoneGameII(piles)
	}
}
