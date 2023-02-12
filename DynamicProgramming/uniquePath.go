package dynamicprogramming

import "fmt"

func uniquePaths(m int, n int) int {

	// Naive solution using DFS
	//paths := gridTravellerDfs(m, n)

	// memoised solution
	dp := make(map[index]int, 0)
	paths := gridTravellerMemo(m, n, dp)

	return paths
}

type index struct {
	row int
	col int
}

// DFS solution with memoisation.
func gridTravellerMemo(row int, col int, dp map[index]int) int {

	if row == 0 || col == 0 {
		return 0
	}

	if row == 1 && col == 1 {
		return 1
	}

	// Add DP optimization
	currentIndex := index{row: row, col: col}

	if value, ok := dp[currentIndex]; ok {

		return value
	}

	down := gridTravellerMemo(row-1, col, dp)
	right := gridTravellerMemo(row, col-1, dp)

	// Memoise the current solution
	dp[currentIndex] = down + right

	return down + right
}

// Naive solution using DFS
func gridTravellerDfs(row int, col int) int {

	if row == 0 || col == 0 {
		return 0
	}

	if row == 1 && col == 1 {
		return 1
	}

	down := gridTravellerDfs(row-1, col)
	right := gridTravellerDfs(row, col-1)

	return down + right
}

func UniquePathCaller() {

	m := 3
	n := 2
	fmt.Print(uniquePaths(m, n))
}
