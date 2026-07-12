package pathexistencequeriesinagraphi

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	id := make([]int, n)
	currentID := 0

	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			currentID++
		}
		id[i] = currentID
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = (id[q[0]] == id[q[1]])
	}
	return res
}
