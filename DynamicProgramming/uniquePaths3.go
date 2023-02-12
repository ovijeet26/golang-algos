package dynamicprogramming

// Leet code link -> https://leetcode.com/problems/unique-paths-iii/
func uniquePathsIII(grid [][]int) int {

	var startX int
	var startY int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 1 {

				startX = i
				startY = j
			}

			if grid[i][j] == 0 {

				emptyStepCount++
			}

		}
	}

	result = 0
	dfs(grid, startX, startY)

	return result
}

var emptyStepCount int
var result int

func dfs(grid [][]int, row int, col int) {

	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {

		return
	}

	if grid[row][col] == -1 {

		return
	}

	if grid[row][col] == 2 {

		if emptyStepCount == 0 {
			result++
			return
		}

	}

	grid[row][col] = -1

	dfs(grid, row+1, col)
	dfs(grid, row-1, col)
	dfs(grid, row, col+1)
	dfs(grid, row, col-1)

	grid[row][col] = 0

}
