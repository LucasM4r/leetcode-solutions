package stonegameiv

import "testing"

func TestWinnerSquareGame(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, false},
		{4, true},
		{7, false},
		{17, false},
	}

	for _, test := range tests {
		result := winnerSquareGame(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %v but got %v", test.n, test.expected, result)
		}
	}
}

func BenchmarkWinnerSquareGame(b *testing.B) {
	for b.Loop() {
		winnerSquareGame(1000)
	}
}
