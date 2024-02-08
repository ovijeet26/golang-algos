package revision

import "fmt"

func threeSum(nums []int) [][]int {
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

func qs(nums []int, left int, right int) {

	if left >= right {
		return
	}

	pivot := nums[(left+right)/2]
	partitionKey := partitionIt(nums, left, right, pivot)
	qs(nums, left, partitionKey-1)
	qs(nums, partitionKey, right)
}

func partitionIt(nums []int, left int, right int, pivot int) int {

	for left <= right {

		for nums[left] < pivot {
			left++
		}

		for nums[right] > pivot {
			right--
		}

		if left <= right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		}
	}

	return left
}

func ThreesumCaller() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
