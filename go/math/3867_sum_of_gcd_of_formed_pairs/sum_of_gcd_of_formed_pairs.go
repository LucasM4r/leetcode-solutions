package sumofgcdofformedpairs

import "sort"

func gcdSum(nums []int) int64 {
	n := len(nums)
	mx := 0

	prefixGcd := make([]int, n)
	var totalSum int64 = 0

	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	for i, value := range nums {
		if value > mx {
			mx = value
		}
		prefixGcd[i] = gcd(value, mx)
	}

	sort.Ints(prefixGcd)

	for i := 0; i < n/2; i++ {
		totalSum += int64(gcd(prefixGcd[i], prefixGcd[n-1-i]))
	}
	return totalSum
}
