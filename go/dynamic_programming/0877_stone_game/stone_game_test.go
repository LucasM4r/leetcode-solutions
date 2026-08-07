package stonegame

import (
	"testing"
)

func TestStoneGame(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  bool
	}{
		{
			name:  "Standard 4-pile case",
			piles: []int{5, 3, 4, 5},
			want:  true,
		},
		{
			name:  "Alternating values 4 piles",
			piles: []int{3, 7, 2, 3},
			want:  true,
		},
		{
			name:  "Larger pile configuration",
			piles: []int{5, 3, 4, 5, 6, 2},
			want:  true,
		},
		{
			name:  "Minimal 2-pile case",
			piles: []int{2, 1},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGame(tt.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStoneGame(b *testing.B) {
	piles := []int{5, 3, 4, 5, 10, 20, 15, 8, 9, 11}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stoneGame(piles)
	}
}
