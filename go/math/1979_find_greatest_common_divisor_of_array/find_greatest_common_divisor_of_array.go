package findgreatestcommondivisorofarray

func findGCD(nums []int) int {
	maxVal := 0
	minVal := 1001

	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}

		if num < minVal {
			minVal = num
		}
	}

	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	return gcd(maxVal, minVal)
}
