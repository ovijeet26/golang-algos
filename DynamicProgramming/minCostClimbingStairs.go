package dynamicprogramming

// Leet code link -> https://leetcode.com/problems/min-cost-climbing-stairs/
func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	dp := make([]int, n)

	for i := 0; i < n; i++ {

		dp[i] = -1
	}
	return min(minStairs(n-1, cost, dp), minStairs(n-2, cost, dp))
}

func minCostClimbingStairsBottomUpTabulation(cost []int) int {

	size := len(cost)

	dp := make([]int, size)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < size; i++ {

		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[size-1], dp[size-2])
}

func optimized_minCostClimbingStairsBottomUpTabulation(cost []int) int {

	size := len(cost)

	for i := 2; i < size; i++ {

		cost[i] = cost[i] + min(cost[i-1], cost[i-2])
	}

	return min(cost[size-1], cost[size-2])
}

func minStairs(current int, cost []int, dp []int) int {

	if current < 0 {

		return 0
	}

	if current == 0 || current == 1 {

		return cost[current]
	}

	if dp[current] != -1 {

		return dp[current]
	}

	dp[current] = cost[current] + min(minStairs(current-1, cost, dp), minStairs(current-2, cost, dp))

	return dp[current]
}

func minStairsBruteForce(currentStep int, cost []int) int {

	if currentStep < 0 {
		return 0
	}

	if currentStep == 0 || currentStep == 1 {
		return cost[currentStep]
	}

	return cost[currentStep] + min(minStairsBruteForce(currentStep-1, cost), minStairsBruteForce(currentStep-2, cost))
}

func min(a int, b int) int {

	if a < b {
		return a
	}

	return b
}
