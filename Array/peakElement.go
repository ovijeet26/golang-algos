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
	right := len(nums) - 1

	for left < right {

		mid := (left + right) / 2

		if isFound(mid, nums) {

			return mid
		}

		if mid != left && nums[mid] <= nums[left] {

			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1

}

func isFound(mid int, nums []int) bool {

	midMinus1 := 0
	midPlus1 := 0

	if mid-1 < 0 {

		midMinus1 = -99999999999
	} else {
		midMinus1 = nums[mid-1]
	}

	if mid+1 >= len(nums) {

		midPlus1 = 99999999999
	} else {
		midPlus1 = nums[mid+1]
	}

	if nums[mid] > midMinus1 && nums[mid] > midPlus1 {

		return true
	}

	return false
}

func PeakElementCaller() {

	// fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{3, 4, 3, 2, 1}))

}
