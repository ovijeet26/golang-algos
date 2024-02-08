package revision

import "fmt"

func quicksort(nums []int, left int, right int) {

	if left >= right {
		return
	}

	pivot := nums[(left+right)/2]
	index := partition(nums, left, right, pivot)
	quicksort(nums, left, index-1)
	quicksort(nums, index, right)
}

func partition(nums []int, left int, right int, pivot int) int {

	for left <= right {

		for nums[left] < pivot {
			left++
		}

		for nums[right] > pivot {
			right--
		}

		if left <= right {

			swap(nums, left, right)
			left++
			right--
		}
	}

	return left
}

func QSCaller() {

	nums := []int{4, 1, 2, -1, 6, 22, 33, 524, -234, 5, 0}

	fmt.Println(nums)
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
