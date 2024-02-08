package revision

func findMin(nums []int) int {

	left, right, minval := 0, len(nums)-1, 9999999999

	for left <= right {

		if nums[left] < nums[right] {
			return minNum(minval, nums[left])
		}

		mid := (left + right) / 2

		minval = minNum(nums[mid], minval)

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return minval
}

func minNum(a int, b int) int {

	if a < b {
		return a
	}

	return b
}
