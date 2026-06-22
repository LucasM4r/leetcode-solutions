package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	length1, length2 := len(nums1), len(nums2)
	total := length1 + length2
	halfTotal := (total + 1) / 2

	low := 0
	high := length1

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := halfTotal - mid1

		left1 := math.MinInt32
		if mid1 > 0 {
			left1 = nums1[mid1-1]
		}

		right1 := math.MaxInt32
		if mid1 < length1 {
			right1 = nums1[mid1]
		}

		left2 := math.MinInt32
		if mid2 > 0 {
			left2 = nums2[mid2-1]
		}

		right2 := math.MaxInt32
		if mid2 < length2 {
			right2 = nums2[mid2]
		}

		if left1 <= right2 && left2 <= right1 {
			if total%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2.0
			}
			return float64(max(left1, left2))
		} else if left1 > right2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}
	return 0.0
}
