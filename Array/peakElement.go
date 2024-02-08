package array

import "fmt"

// Leet code link -> https://leetcode.com/problems/find-peak-element/
func findPeakElement(nums []int) int {

	// Check boundary conditions
	if len(nums) == 0 {

		return -1
	}

	if len(nums) == 1 {

		return 0
	}

	if nums[0] > nums[1] {

		return 0
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {

		return len(nums) - 1
	}

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

	leftValue, rightValue := -9999999999, -9999999999
	if index+1 < len(nums) {

		rightValue = nums[index+1]
	}

	if index-1 >= 0 {

		leftValue = nums[index-1]
	}

	if nums[index] > leftValue && nums[index] > rightValue {

		return true
	}

	return false
}

func PeakElementCaller() {

	// fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{3, 4, 3, 2, 1}))

}
