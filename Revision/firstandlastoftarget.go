package revision

import "fmt"

func searchRange(nums []int, target int) []int {

	index := bs(nums, 0, len(nums)-1, target)
	firstIndex, lastIndex := index, index
	for firstIndex-1 >= 0 && nums[firstIndex-1] == target {
		firstIndex = bs(nums, 0, firstIndex-1, target)
	}

	for lastIndex+1 < len(nums) && nums[lastIndex+1] == target {
		lastIndex = bs(nums, lastIndex+1, len(nums)-1, target)
	}

	return []int{firstIndex, lastIndex}
}

func bs(nums []int, left int, right int, target int) int {

	for left <= right {

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func FnLCaller() {

	nums := []int{2, 2}
	fmt.Print(searchRange(nums, 2))
}
