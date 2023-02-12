package dynamicprogramming

import "fmt"

// Leet code link ->
func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	dp := make(map[int]int, 0)
	minCoin := change(amount, coins, dp)

	if minCoin == 999999999 {
		return -1
	}

	return minCoin - 1
}

func change(amount int, coins []int, dp map[int]int) int {

	if amount == 0 {
		return 1
	}

	if amount < 0 {
		return 0
	}

	if value, ok := dp[amount]; ok {

		return value
	}

	minCoin := 999999999

	for _, coin := range coins {

		remainder := amount - coin
		coins := change(remainder, coins, dp)

		if coins != 0 {
			minCoin = min(minCoin, 1+coins)
		}
	}

	dp[amount] = minCoin
	return minCoin
}

func CoinChangeCaller() {

	fmt.Print(coinChange([]int{2}, 3))
	fmt.Print(coinChange([]int{1, 2, 5}, 11))

}
