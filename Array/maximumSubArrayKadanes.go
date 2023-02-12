package array

// Leet code link -> https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {

	// KADANE'S ALGORITHM
	globalSum := nums[0]
	localSum := nums[0]

	for i := 1; i < len(nums); i++ {

		if localSum < 0 {
			localSum = 0
		}
		localSum += nums[i]

		globalSum = max(globalSum, localSum)
	}

	return globalSum
}

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b
}
