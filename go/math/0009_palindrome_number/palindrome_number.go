package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted_number := 0

	for x > reverted_number {
		reverted_number = reverted_number*10 + x%10
		x /= 10
	}

	return x == reverted_number || x == reverted_number/10
}
