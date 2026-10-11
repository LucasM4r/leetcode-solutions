package stonegameiv

func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false

	for i := 1; i <= n; i++ {
		k := 1
		for k*k <= i {
			if !dp[i-k*k] {
				dp[i] = true
				break
			}
			k++
		}
	}

	return dp[n]
}
