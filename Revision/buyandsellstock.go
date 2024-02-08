package revision

func maxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}
	lowest := prices[0]
	maxProfit := prices[1] - prices[0]

	for i := 1; i < len(prices); i++ {

		difference := prices[i] - lowest

		if prices[i] < lowest {
			lowest = prices[i]
		}

		if difference > maxProfit {
			maxProfit = difference
		}
	}

	if maxProfit < 0 {
		maxProfit = 0
	}
	return maxProfit
}
