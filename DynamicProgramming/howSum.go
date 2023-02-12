package dynamicprogramming

import "fmt"

// howSum() should return an array containing any combination of elements that add up exactly to the
// targetSum. If there are no combinations that add up tot he targetSu, then return null.
// If there are multiple combinations possible, we may return any once.
func howSum(targetSum int, nums []int) []int {

	dp := make(map[int][]int)

	return howSumMemo(nums, targetSum, dp)
	//return howSumRec(nums, targetSum)
}

// Memoised recursive solution
func howSumMemo(nums []int, targetSum int, dp map[int][]int) []int {

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	if arr, ok := dp[targetSum]; ok {

		return arr
	}

	for _, number := range nums {

		remainder := targetSum - number

		result := howSumMemo(nums, remainder, dp)
		//dp[remainder] = result

		if result != nil {

			newArr := append(result, number)
			dp[targetSum] = newArr

			return dp[targetSum]
		}
	}

	dp[targetSum] = nil
	return nil
}

// Naive recursive solution
func howSumRec(nums []int, targetSum int) []int {

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, number := range nums {

		remainder := targetSum - number

		result := howSumRec(nums, remainder)

		if result != nil {

			newArr := append(result, number)
			return newArr
		}
	}

	return nil
}

func HowSumCaller() {

	// nums := []int{7, 14}
	// target := 300

	//nums := []int{2, 3, 5}
	nums := []int{5, 3, 2}
	target := 8

	fmt.Println(howSum(target, nums))
}
