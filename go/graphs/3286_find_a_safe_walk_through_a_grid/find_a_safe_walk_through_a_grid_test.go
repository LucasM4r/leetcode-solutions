package findasafewalkthroughagrid

import "testing"

func TestFindSafeWalk(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		health   int
		expected bool
	}{
		{
			name: "Safe path with no damage",
			grid: [][]int{
				{0, 1, 0, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 0, 1, 0},
			},
			health:   1,
			expected: true,
		},
		{
			name: "Takes some damage but survives",
			grid: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 0},
			},
			health:   4,
			expected: true,
		},
		{
			name: "Not enough health to survive",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			health:   4,
			expected: false,
		},
		{
			name: "Single cell grid - safe",
			grid: [][]int{
				{0},
			},
			health:   1,
			expected: true,
		},
		{
			name: "Single cell grid - dies immediately",
			grid: [][]int{
				{1},
			},
			health:   1,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := findSafeWalk(tc.grid, tc.health)
			if result != tc.expected {
				t.Errorf("Test '%s' failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func BenchmarkFindSafeWalk(b *testing.B) {
	size := 50
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
		for j := range grid[i] {
			grid[i][j] = (i + j) % 2
		}
	}
	health := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findSafeWalk(grid, health)
	}
}
