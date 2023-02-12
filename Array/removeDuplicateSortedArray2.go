package array

import "fmt"

// Leet code link -> https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates2(nums []int) int {

	arraySize := len(nums)
	validIndex := 0
	movingIndex := 1

	for movingIndex < arraySize {

		if nums[validIndex] != nums[movingIndex] {

			validIndex++
			movingIndex++
			continue
		}

		counter := 1
		for movingIndex < arraySize && nums[validIndex] == nums[movingIndex] && counter < 2 {

			validIndex++
			movingIndex++
			counter++
		}

		if movingIndex < arraySize && nums[validIndex] == nums[movingIndex] {

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

func RemoveDuplicate2Caller() {

	//nums := []int{1, 1, 2, 3}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Print(removeDuplicates2(nums))
}
