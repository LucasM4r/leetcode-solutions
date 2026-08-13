package stonegameii

func stoneGameII(piles []int) int {
	n := len(piles)
	suffixSum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i+2*m >= n {
			return suffixSum[i]
		}
		if memo[i][m] != 0 {
			return memo[i][m]
		}
		maxStones := 0
		for x := 1; x <= 2*m; x++ {
			if i+x > n {
				break
			}
			stones := suffixSum[i] - dfs(i+x, max(m, x))
			if stones > maxStones {
				maxStones = stones
			}
		}
		memo[i][m] = maxStones
		return maxStones
	}
	return dfs(0, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
