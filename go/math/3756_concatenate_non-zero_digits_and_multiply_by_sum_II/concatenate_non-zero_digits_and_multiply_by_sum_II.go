package concatenatenonzerodigitsandmultiplybysumii

import "sort"

func sumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1000000007
	if len(s) == 0 || len(queries) == 0 {
		return []int{}
	}
	var index []int
	var values []int

	for i, val := range s {
		if val == '0' {
			continue
		}

		number := int(val - '0')
		index = append(index, i)
		values = append(values, number)
	}

	n := len(values)
	prefSum := make([]int, n+1)
	prefVal := make([]int, n+1)
	pow10 := make([]int, n+1)
	pow10[0] = 1

	for i := 0; i < n; i++ {
		prefSum[i+1] = (prefSum[i] + values[i]) % MOD
		pow10[i+1] = (pow10[i] * 10) % MOD
	}

	for i := 0; i < n; i++ {
		prefVal[i+1] = (prefVal[i]*10 + values[i]) % MOD
	}

	results := make([]int, len(queries))
	for k, q := range queries {
		start, end := sort.SearchInts(index, q[0]), sort.SearchInts(index, q[1]+1)

		if start >= end {
			results[k] = 0
			continue
		}

		sum := (prefSum[end] - prefSum[start] + MOD) % MOD

		dist := end - start
		x := (prefVal[end] - (prefVal[start]*pow10[dist])%MOD + MOD) % MOD

		results[k] = (x * sum) % MOD
	}
	return results
}
