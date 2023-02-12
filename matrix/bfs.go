package matrix

import (
	"fmt"

	queue "github.com/ovijeet/go/algos/Queue"
)

type indices struct {
	row int
	col int
}

func bfs(matrix [][]int, row int, col int, visited [][]bool) []int {

	value := []int{}
	toVisit := queue.Queue{}

	toVisit.Enqueue(indices{row: row, col: col})

	for toVisit.Size() != 0 {

		currentValue := toVisit.Dequeue().(indices)
		currentRow := currentValue.row
		currentCol := currentValue.col

		if currentRow < 0 || currentRow > len(matrix)-1 ||
			currentCol < 0 || currentCol > len(matrix[0])-1 ||
			visited[currentRow][currentCol] {

			continue
		}

		value = append(value, matrix[currentRow][currentCol])
		visited[currentRow][currentCol] = true

		up := indices{row: currentRow - 1, col: currentCol}
		right := indices{row: currentRow, col: currentCol + 1}
		down := indices{row: currentRow + 1, col: currentCol}
		left := indices{row: currentRow, col: currentCol - 1}

		toVisit.Enqueue(up)
		toVisit.Enqueue(right)
		toVisit.Enqueue(down)
		toVisit.Enqueue(left)
	}

	return value
}

func BfsCaller() {

	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	visited := [][]bool{}

	for i := 0; i < len(input); i++ {
		tempVisited := []bool{false, false, false, false, false}
		visited = append(visited, tempVisited)
	}

	fmt.Print(bfs(input, 2, 2, visited))
}
