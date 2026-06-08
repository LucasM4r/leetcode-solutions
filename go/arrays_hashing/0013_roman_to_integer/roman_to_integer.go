package romantointeger

func getRomanValue(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

// romanToInt finds the integer value of a Roman numeral string by iterating through it and comparing adjacent characters.
// Time Complexity: O(N)
// Space Complexity: O(1)
func romanToInt(s string) int {
	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentVal := getRomanValue(s[i])

		if i < n-1 && currentVal < getRomanValue(s[i+1]) {
			total -= currentVal
		} else {
			total += currentVal
		}
	}

	return total
}
