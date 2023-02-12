package matrix

import "fmt"

// Leet code link -> https://leetcode.com/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {

	maxIslandArea := 0

	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	rowLength := len(grid)
	colLength := len(grid[0])

	// Scan the matrix to check for a land
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			if grid[i][j] == 1 {
				area = 0
				dfsAreaIsland(grid, i, j, directions)
				maxIslandArea = max(maxIslandArea, area)
			}
		}
	}

	return maxIslandArea
}

var area int

func dfsAreaIsland(grid [][]int, row int, col int, directions [][]int) {

	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == 0 {

		return
	}

	// Since it's already visisted, we can mark it as water (0)
	grid[row][col] = 0
	area++

	for _, direction := range directions {

		rowOffset := direction[0]
		colOffset := direction[1]

		dfsAreaIsland(grid, row+rowOffset, col+colOffset, directions)
	}

}

func max(a int, b int) int {

	if a > b {

		return a
	}

	return b
}

func AreaIslandCaller() {

	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}
