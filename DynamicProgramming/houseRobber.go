package dynamicprogramming

// Leet code link -> https://www.youtube.com/watch?v=73r3KWiEvyk&ab_channel=NeetCode
func rob(nums []int) int {

	if len(nums) == 1 {

		return nums[0]
	}

	rob0 := 0
	rob1 := 0

	// We make our array like this:
	// [ ...., rob0, rob1, n, n+1, ..... ]
	for i := 0; i < len(nums); i++ {

		// At each position, we either take costa t that position and max plus amx or we ignore that and take max at adjascent position.
		costAtPosition := max(nums[i]+rob0, rob1)
		rob0 = rob1
		rob1 = costAtPosition
	}

	return max(rob0, rob1)
}

func max(a int, b int) int {

	if a > b {

		return a
	}

	return b
}

// Naive non-optimized solution
func robNaiveDp(nums []int) int {

	if len(nums) == 1 {

		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < len(nums); i++ {

		dp[i] = nums[i] + maxTillIndex(dp, i-2)
	}

	return maxTillIndex(dp, len(nums)-1)
}

func maxTillIndex(arr []int, index int) int {

	max := arr[0]

	for i := 0; i <= index; i++ {

		if arr[i] > max {

			max = arr[i]
		}
	}

	return max
}
