package removeelement

func removeElement(nums []int, val int) int {
	arrayPointer := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[arrayPointer] = nums[i]

			arrayPointer++
		}
	}
	return arrayPointer
}
