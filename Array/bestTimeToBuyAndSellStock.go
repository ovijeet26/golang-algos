package array

// Leet code link -> https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {

	// Employ two pointer technique to check possible buy and sell dates.
	buy := 0
	sell := 0

	maxProfit := 0
	for sell = 1; sell < len(prices); sell++ {

		profit := prices[sell] - prices[buy]
		maxProfit = max(maxProfit, profit)

		// If the payout is not profitable or breakeven, we change the buying day
		if profit <= 0 {

			buy = sell
		}
	}

	return maxProfit
}
