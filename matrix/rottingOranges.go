package matrix

import (
	"fmt"

	queue "github.com/ovijeet/go/algos/Queue"
)

// Leet code link -> https://leetcode.com/problems/rotting-oranges/
func orangesRotting(grid [][]int) int {

	rowLength := len(grid)
	colLength := len(grid[0])

	if rowLength == 0 {

		return 0
	}

	minutes := 0
	freshOranges := 0
	rottenOranges := []indices{}

	directions := [][]int{
		{1, 0},  // up
		{0, 1},  // right
		{-1, 0}, // down
		{0, -1}, // left
	}

	// Compute the position of rotten oranges and the count of fresh oranges.
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			if grid[i][j] == 2 {

				rottenOranges = append(rottenOranges, indices{row: i, col: j})
			}

			if grid[i][j] == 1 {

				freshOranges++
			}
		}
	}

	// If no fresh oranges are present, then no need to process.
	if freshOranges == 0 {
		return 0
	}

	// If no rotten oranges are present, then no need to process.
	if freshOranges == 0 {
		return -1
	}

	toVisit := queue.Queue{}

	// Populate the queue with all the rotten oranges
	for _, rottenOrange := range rottenOranges {

		toVisit.Enqueue(rottenOrange)
	}

	levelCounter := toVisit.Size()

	for toVisit.Size() > 0 {

		currentOrange := toVisit.Dequeue().(indices)
		currentRow := currentOrange.row
		currentCol := currentOrange.col

		levelCounter--

		for _, direction := range directions {

			nextRow := currentRow + direction[0]
			nextCol := currentCol + direction[1]

			if nextRow < 0 || nextRow >= rowLength || nextCol < 0 || nextCol >= colLength || grid[nextRow][nextCol] != 1 {

				continue
			}

			toVisit.Enqueue(indices{row: nextRow, col: nextCol})
			// Mark next valid index as rotten
			grid[nextRow][nextCol] = 2
			// Decrease the count of fresh oranges
			freshOranges--

		}

		if levelCounter == 0 {

			minutes++
			// Stop processing if all the fresh oranges are rotten.
			if freshOranges == 0 {
				break
			}
			levelCounter = toVisit.Size()
		}
	}

	// If there are still fresh oranges left, that means we couldn’t rot all of the oranges
	if freshOranges > 0 {

		return -1
	}

	return minutes
}

// type indices struct {

// 	row int
// 	col int
// }

// // Regular queue implementation

// type Queue struct {
// 	items []interface{}
// }

// func (q *Queue) Enqueue(data interface{}) {

// 	q.items = append(q.items, data)
// }

// func (q *Queue) Dequeue() interface{} {

// 	toRemove := q.items[0]

// 	// Remove first element from slice
// 	q.items = q.items[1:]
// 	return toRemove
// }

// func (q *Queue) Size() int {

// 	return len(q.items)
// }

func RottingOrangeCaller() {

	// grid := [][]int{
	// 	{2, 1, 1},
	// 	{1, 1, 0},
	// 	{0, 1, 1},
	// }

	grid2 := [][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2},
	}

	mins := orangesRotting(grid2)

	fmt.Print(mins)
}
