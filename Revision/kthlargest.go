package revision

func findKthLargest(nums []int, k int) int {

	idxToFind := len(nums) - k
	quickselect(nums, 0, len(nums)-1, idxToFind)

	return nums[idxToFind]
}

func quickselect(nums []int, left int, right int, idxToFind int) {

	for left < right {

		pivot := nums[right]
		index := qspartition(nums, left, right, pivot)

		if index == idxToFind {
			return
		} else if index < idxToFind {
			quickselect(nums, index+1, right, idxToFind)
		} else {
			quickselect(nums, left, index-1, idxToFind)
		}
	}
}

func qspartition(nums []int, left int, right int, pivot int) int {

	for left <= right {
		for nums[left] < pivot {
			left++
		}

		for nums[right] > pivot {
			right--
		}

		if left <= right {
			xswap(nums, left, right)
			left--
			right++
		}
	}

	return left
}

func xswap(nums []int, i int, j int) {

	nums[i], nums[j] = nums[j], nums[i]
}
