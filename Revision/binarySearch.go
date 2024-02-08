package revision

import "fmt"

func binarySearch(nums []int, searchKey int) int {

	left, right := 0, len(nums)-1

	for left <= right {

		mid := (left + right) / 2

		if nums[mid] == searchKey {
			return mid
		} else if nums[mid] < searchKey {
			left = mid + 1
			continue
		} else {
			right = mid - 1
			continue
		}
	}

	return -1
}

func BSCaller() {

	nums := []int{1, 2, 3, 4, 5, 6, 77, 345, 2333, 44444}

	fmt.Print(binarySearch(nums, 4444))
}
