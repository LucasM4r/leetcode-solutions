package twosum

// TwoSumHashMap finds the indices of two numbers using a HashMap and add up to the target.
// Time Complexity: O(N)
// Space Complexity: O(N)
func TwoSumHashMap(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := hash[complement]; ok {
			return []int{index, i}
		}

		hash[num] = i
	}

	return nil
}

// TwoSumArray finds the indices of two numbers using a Array Brute Force and add up to the target.
// Time Complexity: O(N^2)
// Space Complexity: O(1)
func TwoSumArray(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
