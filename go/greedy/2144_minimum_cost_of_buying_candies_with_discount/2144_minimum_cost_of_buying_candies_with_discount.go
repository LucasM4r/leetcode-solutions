package minimumcostofbuyingcandieswithdiscount

import "sort"

func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	var totalCost int

	for i := 0; i < len(cost); i++ {
		if (i+1)%3 != 0 {
			totalCost += cost[i]
		}
	}

	return totalCost
}
