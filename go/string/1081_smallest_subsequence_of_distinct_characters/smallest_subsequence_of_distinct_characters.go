package smallestsubsequenceofdistinctcharacters

func smallestSubsequence(s string) string {
	lastOccurence := [26]int{}
	visited := [26]bool{}

	for i, char := range s {
		lastOccurence[char-'a'] = i
	}

	stack := []byte{}

	for i, char := range s {
		if visited[char-'a'] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > byte(char) && lastOccurence[stack[len(stack)-1]-'a'] > i {
			visited[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, byte(char))
		visited[char-'a'] = true
	}

	return string(stack)
}
