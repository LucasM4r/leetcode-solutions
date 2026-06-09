package stringtointeger

import "math"

// myAtoi converts a string to a 32-bit signed integer.
// Time Complexity: O(N), where N is the length of the string 's', as it iterates through the string at most once.
// Space Complexity: O(1), as it only uses a few variables for pointers and calculations, requiring no extra scaling memory.
func myAtoi(s string) int {
	i := 0
	n := len(s)

	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	sign := 1
	switch s[i] {
	case '-':
		sign = -1
		i++
	case '+':
		i++
	}

	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digitValue := int(s[i] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digitValue > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = (result * 10) + digitValue
		i++
	}

	return result * sign
}
