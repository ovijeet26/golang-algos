package matrix

import (
	"fmt"

	queue "github.com/ovijeet/go/algos/Queue"
)

func NumIslandsCaller() {

	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}

	fmt.Print(numIslands(grid))
	//fmt.Print(numIslandsVisited(grid))
}

// leet code link -> https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {

	islandCount := 0

	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	toVisit := queue.Queue{}

	rowLength := len(grid)
	// Considering that the matrix is n x m
	colLength := len(grid[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			// Do selective BFS till we find all adjacent 1s. Meanwhile, we have to make the visited 1s as 0.
			if grid[i][j] == '1' {

				islandCount++
				toVisit.Enqueue(indices{row: i, col: j})

				for toVisit.Size() > 0 {
					currentNode := toVisit.Dequeue().(indices)
					currentRow := currentNode.row
					currentCol := currentNode.col
					// Mark this as visited, hence changing to 0.
					grid[currentRow][currentCol] = '0'

					// Iterate over all directions for BFS.
					for _, value := range directions {

						nextRow := currentRow + value[0]
						nextCol := currentCol + value[1]

						// Ignore the boundary conditions
						if nextRow < 0 || nextRow >= rowLength || nextCol < 0 || nextCol >= colLength {

							continue
						}

						if grid[nextRow][nextCol] == '1' {
							toVisit.Enqueue(indices{row: nextRow, col: nextCol})
							grid[nextRow][nextCol] = '0'
						}
					}

				}
			}
		}
	}

	return islandCount
}

//..........

func numIslandsDfs(grid [][]byte) int {

	islandCount := 0

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

			if grid[i][j] == '1' {
				dfsIsland(grid, i, j, directions)
				islandCount++
			}
		}
	}

	return islandCount
}

func dfsIsland(grid [][]byte, row int, col int, directions [][]int) {

	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {

		return
	}

	// Since it's already visisted, we can mark it as water (0)
	grid[row][col] = '0'

	for _, direction := range directions {

		rowOffset := direction[0]
		colOffset := direction[1]

		dfsIsland(grid, row+rowOffset, col+colOffset, directions)
	}

}

//..........
func numIslandsVisited(grid [][]byte) int {

	islandCount := 0

	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	toVisit := queue.Queue{}

	rowLength := len(grid)
	// Considering that the matrix is n x m
	colLength := len(grid[0])

	visited := [][]bool{}

	for i := 0; i < rowLength; i++ {
		tempVisited := []bool{}

		for i := 0; i < colLength; i++ {

			tempVisited = append(tempVisited, false)
		}
		visited = append(visited, tempVisited)
	}

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			// Do selective BFS till we find all adjacent 1s. Meanwhile, we have to make the visited 1s as 0.
			if grid[i][j] == '1' && !visited[i][j] {

				islandCount++
				toVisit.Enqueue(indices{row: i, col: j})

				for toVisit.Size() > 0 {
					currentNode := toVisit.Dequeue().(indices)
					currentRow := currentNode.row
					currentCol := currentNode.col

					// Mark this as visited
					visited[currentRow][currentCol] = true

					// Iterate over all directions for BFS.
					for _, value := range directions {

						nextRow := currentRow + value[0]
						nextCol := currentCol + value[1]

						// Ignore the boundary conditions
						if nextRow < 0 || nextRow >= rowLength || nextCol < 0 ||
							nextCol >= colLength || visited[nextRow][nextCol] {

							continue
						}

						toVisit.Enqueue(indices{row: nextRow, col: nextCol})
					}

				}
			}
		}
	}

	return islandCount
}

// type indices struct {
// 	row int
// 	col int
// }
