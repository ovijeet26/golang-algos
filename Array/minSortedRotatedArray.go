package array

// Leet code link -> https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {

	// We will deploy a modified binary search.
	left := 0
	right := len(nums) - 1
	minVal := 999999999999999

	for left <= right {

		// If the sub array is already sorted, we simply check min with the left-most element.
		if nums[left] < nums[right] {

			minVal = min(minVal, nums[left])
			break
		}

		mid := (left + right) / 2

		// Check if mid is lesser than minimum value
		minVal = min(minVal, nums[mid])

		// If mid is greater than or equal to left, then check right subarray
		if nums[mid] >= nums[left] {

			left = mid + 1
		} else {

			// Else check left sub array.
			right = mid - 1
		}
	}

	return minVal
}

func min(a int, b int) int {

	if a < b {

		return a
	}

	return b
}
