package numberofuniquexortripletsi

import "math/bits"

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	return 1 << (bits.Len(uint(n)))
}
