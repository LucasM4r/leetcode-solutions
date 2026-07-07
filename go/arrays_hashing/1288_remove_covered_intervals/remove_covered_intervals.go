package removecoveredintervals

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	end := 0

	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] <= end {
			continue
		}

		end = intervals[i][1]
		count++
	}
	return count
}
