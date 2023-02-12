package dynamicprogramming

// Leet code link -> https://leetcode.com/problems/house-robber-ii/
func rob2(nums []int) int {

	size := len(nums)

	if size == 1 {
		return nums[0]
	}

	// Calculate the max cost excluding the first house
	maxCostExcludingHouse1 := houseRobber(nums[1:])
	// Calculate the max cost excluding the last house
	maxCostExcludingHouseN := houseRobber(nums[:size-1])

	return max(maxCostExcludingHouse1, maxCostExcludingHouseN)
}

// See house robber solution for details
func houseRobber(nums []int) int {

	size := len(nums)

	rob0 := 0
	rob1 := 0

	for i := 0; i < size; i++ {

		costAtPosition := max(nums[i]+rob0, rob1)
		rob0 = rob1
		rob1 = costAtPosition
	}

	return max(rob0, rob1)
}
