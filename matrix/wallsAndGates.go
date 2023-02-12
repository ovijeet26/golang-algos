package matrix

import (
	"fmt"

	queue "github.com/ovijeet/go/algos/Queue"
)

// Leet code link -> https://leetcode.com/problems/walls-and-gates/
// You are given a m x n 2D grid initialized with these three possible values.
//
// -1 - A wall or an obstacle. 0 - A gate. INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647. Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.
//
// For example, given the 2D grid:
//
// INF  -1  0  INF
// INF INF INF  -1
// INF  -1 INF  -1
//   0  -1 INF INF
// After running your function, the 2D grid should be:
//
//   3  -1   0   1
//   2   2   1  -1
//   1  -1   2  -1
//   0  -1   3   4

func wallsAndGates(grid [][]int) {

	if len(grid) == 0 {
		return
	}

	gates := []indices{}

	rowLength := len(grid)
	colLength := len(grid[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			if grid[i][j] == 0 {

				gates = append(gates, indices{row: i, col: j})
			}
		}
	}

	if len(gates) == 0 {
		return
	}

	directions := [][]int{
		{1, 0},  // up
		{0, 1},  // right
		{-1, 0}, // down
		{0, -1}, // left
	}
	INF := 8

	toVisit := queue.Queue{}

	for _, gate := range gates {

		toVisit.Enqueue(gate)
	}

	steps := 1
	levelCounter := toVisit.Size()

	for toVisit.Size() > 0 {

		currentNode := toVisit.Dequeue().(indices)
		currentRow := currentNode.row
		currrentCol := currentNode.col
		levelCounter--
		for _, direction := range directions {

			nextRow := currentRow + direction[0]
			nextCol := currrentCol + direction[1]

			if nextRow < 0 || nextRow >= rowLength || nextCol < 0 || nextCol >= colLength {

				continue
			}

			if grid[nextRow][nextCol] == INF {
				toVisit.Enqueue(indices{row: nextRow, col: nextCol})
				grid[nextRow][nextCol] = steps
			}
		}

		if levelCounter == 0 {

			steps++
			levelCounter = toVisit.Size()
		}
	}
}

func wallsAndGatesDfs(grid [][]int) {

	if len(grid) == 0 {
		return
	}

	gates := []indices{}

	rowLength := len(grid)
	colLength := len(grid[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			if grid[i][j] == 0 {

				gates = append(gates, indices{row: i, col: j})
			}
		}
	}

	if len(gates) == 0 {
		return
	}

	directions := [][]int{
		{1, 0},  // up
		{0, 1},  // right
		{-1, 0}, // down
		{0, -1}, // left
	}
	//INF := 8

	for _, gate := range gates {

		wallDfs(grid, directions, gate.row, gate.col, 0)
	}
}

func wallDfs(grid [][]int, directions [][]int, row int, col int, step int) {

	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] > step {

		return
	}

	grid[row][col] = step

	for _, direction := range directions {

		wallDfs(grid, directions, row+direction[0], col+direction[1], step+1)
	}
}

// type indices struct {

// 	row int
// 	col int
// }

func WallsAndGateCaller() {

	//INF := int(math.Inf(1))
	INF := 8

	grid := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}

	//wallsAndGates(grid)
	wallsAndGatesDfs(grid)
	fmt.Println(grid)
}
