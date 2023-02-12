package dynamicprogramming

import "fmt"

// Write a function, which should return an array containing the shortest combination
// of numbers that add up exactly to the targetSum.
// If there is a tie for the shortest combination, we may return any one.
func bestSum(targetSum int, nums []int) []int {

	dp := make(map[int][]int, 0)

	return bestSumMemo(nums, targetSum, dp)
	//return bestSumRec(nums, targetSum)
}

// Memoised recursive solution
func bestSumMemo(nums []int, targetSum int, dp map[int][]int) []int {

	if targetSum == 0 {

		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	if value, ok := dp[targetSum]; ok {

		return value
	}

	// shortestCombo := []int{}
	// shortestCombo = nil
	var shortestCombo []int

	for _, num := range nums {

		remainder := targetSum - num

		combo := bestSumMemo(nums, remainder, dp)

		if combo != nil {

			combo = append(combo, num)

			if shortestCombo == nil || len(combo) < len(shortestCombo) {

				shortestCombo = combo
			}
		}
	}

	dp[targetSum] = shortestCombo

	return shortestCombo
}

// Naive recursive solution
func bestSumRec(nums []int, targetSum int) []int {

	if targetSum == 0 {

		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	shortestCombo := []int{}
	shortestCombo = nil

	for _, num := range nums {

		remainder := targetSum - num

		combo := bestSumRec(nums, remainder)

		if combo != nil {

			combo = append(combo, num)

			if shortestCombo == nil || len(combo) < len(shortestCombo) {

				shortestCombo = combo
			}
		}
	}

	return shortestCombo
}

func BestSumCaller() {

	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	fmt.Println(bestSum(8, []int{1, 4, 5}))
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}))
}
