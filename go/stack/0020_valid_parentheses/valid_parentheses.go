package validparentheses

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		input := s[i]

		if input == '(' || input == '{' || input == '[' {
			stack = append(stack, input)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			xor := top ^ input
			if xor != 1 && xor != 6 {
				return false
			}
		}
	}
	return len(stack) == 0
}
