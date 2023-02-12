package array

import "fmt"

// Leet code link -> https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {

	arraySize := len(nums)
	validIndex := 0
	movingIndex := 1

	for movingIndex < arraySize {

		if nums[validIndex] != nums[movingIndex] {

			validIndex++
			movingIndex++
			continue
		}

		if nums[validIndex] == nums[movingIndex] {

			for movingIndex < arraySize && nums[validIndex] == nums[movingIndex] {

				movingIndex++
			}

			if movingIndex >= arraySize {
				break
			}

			validIndex++
			nums[validIndex] = nums[movingIndex]

		}
	}

	return validIndex + 1
}

func RemoveDuplicateCaller() {

	nums := []int{1, 1, 2, 3}
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Print(removeDuplicates(nums))
}
