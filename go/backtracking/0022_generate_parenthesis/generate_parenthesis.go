package generateparenthesis

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(int, int, string)

	backtrack = func(open, close int, actualString string) {
		if len(actualString) == n*2 {
			result = append(result, actualString)
			return
		}

		if open < n {
			backtrack(open+1, close, actualString+"(")
		}

		if close < open {
			backtrack(open, close+1, actualString+")")
		}
	}

	backtrack(0, 0, "")
	return result
}
