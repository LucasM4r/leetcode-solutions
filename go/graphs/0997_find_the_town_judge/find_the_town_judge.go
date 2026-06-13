package findthetownjudge

func findJudge(n int, trust [][]int) int {
	scores := make([]int, n+1)

	for i := 0; i < len(trust); i++ {
		scores[trust[i][0]]--
		scores[trust[i][1]]++
	}

	for i := 1; i < len(scores); i++ {
		if scores[i] == n-1 {
			return i
		}
	}

	return -1
}
