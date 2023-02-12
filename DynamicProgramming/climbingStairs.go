package dynamicprogramming

// Leet code link -> https://leetcode.com/problems/climbing-stairs/submissions/
func climbStairs(n int) int {

	if n <= 2 {
		return n
	}

	// We are opting a bottom-up approach.
	// Initially, the number of ways to reach step n from n = 1 and
	// the number of ways to reach step n from n-1 = 1
	nPlusOne := 1
	nPlusTwo := 1

	stepCount := 0

	// Starting iteration from n-2, and computing till 0 -> our starting point.
	for i := n - 2; i >= 0; i-- {

		// Calculate steps required to reach n from current step
		stepCount = nPlusOne + nPlusTwo

		nPlusTwo = nPlusOne
		nPlusOne = stepCount
	}

	return stepCount
}
