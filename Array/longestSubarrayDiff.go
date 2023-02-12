package array

import "fmt"

// Leet code link -> https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func longestSubarray(nums []int, limit int) int {

	leftIndex, rightIndex := 0, 0

	maxWindowSize := 0
	maxDeq := Deque{}
	minDeq := Deque{}

	// Iterate over the array and increase the sliding window
	for rightIndex = 0; rightIndex < len(nums); rightIndex++ {

		// Pop all elements in the right of the max Deque which are smaller than the current item
		for maxDeq.Size() > 0 && nums[rightIndex] > nums[maxDeq.PeekRight()] {

			maxDeq.PopRight()
		}

		// Pop all elements in the right of the min Deque which are greater than the current item
		for minDeq.Size() > 0 && nums[rightIndex] < nums[minDeq.PeekRight()] {

			minDeq.PopRight()
		}

		//1, 5, 6, 7, 8, 10, 6, 5, 6
		// Add the current element to the Max and Min Deque
		maxDeq.Push(rightIndex)
		minDeq.Push(rightIndex)

		// Compute the absolute difference
		absoluteDifference := abs(nums[maxDeq.PeekLeft()] - nums[minDeq.PeekLeft()])

		// Update boundaries in case we have a violation.
		if absoluteDifference > limit {

			if leftIndex == maxDeq.PeekLeft() {

				maxDeq.PopLeft()
			}

			if leftIndex == minDeq.PeekLeft() {

				minDeq.PopLeft()
			}

			leftIndex++
			continue
		}

		// Compute the max window size in case of no violations.
		maxWindowSize = max(maxWindowSize, (rightIndex-leftIndex)+1)
	}

	return maxWindowSize
}

func abs(a int) int {

	if a < 0 {
		return -a
	}

	return a
}

func SubArrayDiffCaller() {

	// nums := []int{4, 2, 2, 2, 4, 4, 2, 2}
	// limit := 0

	nums := []int{1, 5, 6, 7, 8, 10, 6, 5, 6}
	limit := 4
	fmt.Print(longestSubarray(nums, limit))
}
