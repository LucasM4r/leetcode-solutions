package numberofpathswithmaxscore

func pathsWithMaxScore(board []string) []int {
	type Cell struct {
		score int
		paths int
	}

	n := len(board)
	dp := make([][]Cell, n)
	for i := range dp {
		dp[i] = make([]Cell, n)
	}

	dp[n-1][n-1] = Cell{score: 0, paths: 1}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			char := board[i][j]
			if char == 'X' || char == 'S' {
				continue
			}

			val := 0
			if char != 'E' {
				val = int(char - '0')
			}

			maxNeighborScore := -1
			neighborPaths := 0

			directions := [][2]int{{i + 1, j}, {i, j + 1}, {i + 1, j + 1}}

			for _, dir := range directions {
				ni, nj := dir[0], dir[1]
				if ni < n && nj < n {
					neighbor := dp[ni][nj]
					if neighbor.score > maxNeighborScore {
						maxNeighborScore = neighbor.score
						neighborPaths = neighbor.paths
					} else if neighbor.score == maxNeighborScore {
						neighborPaths = (neighborPaths + neighbor.paths) % 1000000007
					}
				}
			}
			if maxNeighborScore != -1 {
				dp[i][j].score = maxNeighborScore + val
				dp[i][j].paths = neighborPaths
			}
		}
	}
	if dp[0][0].paths == 0 {
		return []int{0, 0}
	}
	return []int{dp[0][0].score, dp[0][0].paths}
}
