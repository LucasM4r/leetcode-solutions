package ranktransformofanarray

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	rankMap := make(map[int]int)
	rank := 1
	for i, val := range sorted {
		if i > 0 && val != sorted[i-1] {
			rank++
		}
		if _, exists := rankMap[val]; !exists {
			rankMap[val] = rank
		}
	}

	result := make([]int, len(arr))
	for i, val := range arr {
		result[i] = rankMap[val]
	}

	return result
}
