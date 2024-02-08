package array

import "fmt"

// Leet code link -> https://leetcode.com/problems/3sum/
func threeSum_Cleaner(nums []int) [][]int {
	// Sorting the array
	qs(nums, 0, len(nums)-1)
	pointer, resultSet := 0, make([][]int, 0)

	// Loop for checking against 1 value.
	for pointer < len(nums) {
		if pointer > 0 && nums[pointer-1] == nums[pointer] {
			pointer++
			continue
		}

		left, right := pointer+1, len(nums)-1
		// Two sum implememtation
		// Two pointer on both sides,
		// Increase left if Sum is lesser than the target since the array is sorted and adding left will only reduce the total sum
		// Same as left we have to decrease the right when sum is greater than the target.
		for left < right {
			sumOfThree := nums[pointer] + nums[left] + nums[right]
			if sumOfThree > 0 {
				right--
			} else if sumOfThree < 0 {
				left++
			} else {
				resultSet = append(resultSet, []int{nums[pointer], nums[left], nums[right]})
				left++
				right--
				// If the previous value is same.
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
		pointer++
	}
	return resultSet
}

func threeSum(nums []int) [][]int {

	// Sort the entire array
	quickSort(nums)
	result := make([][]int, 0)
	size := len(nums)

	left := 0
	right := size - 1

	for i := 0; i < size; i++ {

		// If we come across the same number twice, we will skip it to avoid duplicates
		if i > 0 && nums[i] == nums[i-1] {

			continue
		}

		left = i + 1
		right = size - 1
		fixedValueIndex := i

		for left < right {

			if uniquelyAddedToZero(nums, fixedValueIndex, left, right) {

				result = append(result, []int{nums[fixedValueIndex], nums[left], nums[right]})
			}

			if nums[fixedValueIndex]+nums[left]+nums[right] > 0 {

				right--
				// If we come across the same right again, we will skip it to avoid duplicates
				for nums[right] == nums[right+1] {
					right--
				}
			} else {

				left++
				// If we come across the same left again, we will skip it to avoid duplicates
				for nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return result
}

func uniquelyAddedToZero(nums []int, i int, j int, k int) bool {

	if i == j || j == k || i == k {

		return false
	}

	if nums[i]+nums[j]+nums[k] == 0 {
		return true
	}

	return false
}
func quickSort(nums []int) {

	qs(nums, 0, len(nums)-1)
}

func qs(nums []int, left int, right int) {

	if left < right {

		partitionIndex := partition(nums, left, right)
		qs(nums, left, partitionIndex-1)
		qs(nums, partitionIndex+1, right)
	}
}

func partition(nums []int, left int, right int) int {

	pivot := nums[right]
	partitionIndex := left

	for i := left; i < right; i++ {

		if nums[i] <= pivot {

			swap(nums, i, partitionIndex)
			partitionIndex++
		}
	}

	swap(nums, partitionIndex, right)

	return partitionIndex
}

func swap(nums []int, i int, j int) {

	nums[i], nums[j] = nums[j], nums[i]
}

func ThreeSumCaller() {

	//nums := []int{-1, 0, 1, 2, -1, -4}

	nums := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}

	fmt.Println(threeSum(nums))
}
