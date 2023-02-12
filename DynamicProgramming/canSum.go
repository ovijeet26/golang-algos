package dynamicprogramming

import "fmt"

// Given an array of integers, return true if a target value can be formed by adding them.
func canSum(target int, nums []int) bool {

	dp := make(map[int]bool, 0)

	return canSumMemo(nums, target, dp)
	//return canSumRec(nums, target)
}

// Memoised recursive solution
func canSumMemo(nums []int, targetSum int, dp map[int]bool) bool {

	if targetSum == 0 {

		return true
	}

	if targetSum < 0 {

		return false
	}

	if value, ok := dp[targetSum]; ok {

		return value
	}

	for _, num := range nums {

		remainder := targetSum - num
		dp[remainder] = canSumMemo(nums, remainder, dp)

		if dp[remainder] == true {
			return true
		}
	}

	dp[targetSum] = false
	return false
}

// Naaive recursive solution
func canSumRec(nums []int, targetSum int) bool {

	if targetSum == 0 {

		return true
	}

	if targetSum < 0 {

		return false
	}

	for _, num := range nums {

		remainder := targetSum - num
		if canSumRec(nums, remainder) {
			return true
		}
	}

	return false
}

func CanSumCaller() {

	// nums := []int{2, 3}
	// target := 7

	// nums := []int{2, 3, 5}
	// target := 8

	nums := []int{7, 14}
	target := 300

	fmt.Print(canSum(target, nums))
}
