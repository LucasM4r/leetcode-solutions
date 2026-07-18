package reverseinteger

import "math"

func reverse(x int) int {
	reversedInteger := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		if reversedInteger > math.MaxInt32/10 || (reversedInteger == math.MaxInt32/10 && digit > 7) {
			return 0
		}

		if reversedInteger < math.MinInt32/10 || (reversedInteger == math.MinInt32/10 && digit < -8) {
			return 0
		}

		reversedInteger = reversedInteger*10 + digit
	}
	return reversedInteger
}
