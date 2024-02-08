package revision

func findPeakElement(nums []int) int {

	left := 0
	right := len(nums)

	for left <= right {

		mid := (left + right) / 2

		if isPeak(nums, mid) {
			return mid
		}

		if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func isPeak(nums []int, index int) bool {

	left, right := -9999999999, -9999999999
	if index+1 < len(nums) {

		right = nums[index+1]
	}

	if index-1 >= 0 {

		left = nums[index-1]
	}

	if nums[index] > left && nums[index] > right {

		return true
	}

	return false
}
