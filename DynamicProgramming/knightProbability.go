package dynamicprogramming

import "fmt"

// Leet code link -> https://leetcode.com/problems/knight-probability-in-chessboard/
func knightProbability(n int, k int, row int, column int) float64 {

	knightDirections := [][]int{
		{-2, -1},
		{-2, 1},
		{-1, 2},
		{1, 2},
		{2, 1},
		{2, -1},
		{1, -2},
		{-1, -2},
	}

	dp := make(map[string]float64)

	return knightPMemoised(k, n, row, column, knightDirections, dp)
}

func knightPMemoised(k int, size int, row int, col int, directions [][]int, dp map[string]float64) float64 {

	if row < 0 || row >= size || col < 0 || col >= size {

		return 0
	}

	if k == 0 {

		return 1
	}

	keyString := fmt.Sprintf("%v%v%v", k, row, col)

	if value, ok := dp[keyString]; ok {

		return value
	}

	finalResult := 0.0000
	for _, direction := range directions {

		rowOffset := direction[0]
		colOffset := direction[1]
		finalResult += knightPMemoised(k-1, size, row+rowOffset, col+colOffset, directions, dp) / 8
	}

	dp[keyString] = finalResult
	return finalResult
}

func knightPRecursive(k int, size int, row int, col int, directions [][]int) float64 {

	if row < 0 || row >= size || col < 0 || col >= size {

		return 0
	}

	if k == 0 {

		return 1
	}

	var finalResult float64
	for _, direction := range directions {

		rowOffset := direction[0]
		colOffset := direction[1]
		finalResult += knightPRecursive(k-1, size, row+rowOffset, col+colOffset, directions) / 8
	}

	return finalResult
}
