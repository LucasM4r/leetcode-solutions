package stonegameiii

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	db := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		maxVal := math.MinInt32
		sum := 0

		for k := 1; k <= 3 && i+k-1 < n; k++ {
			sum += stoneValue[i+k-1]
			currentScore := sum - db[i+k]
			if currentScore > maxVal {
				maxVal = currentScore
			}
		}
		db[i] = maxVal
	}

	if db[0] > 0 {
		return "Alice"
	} else if db[0] < 0 {
		return "Bob"
	}
	return "Tie"
}
