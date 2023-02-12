package matrix

import "fmt"

// Please note that when you pass a slice to a function,
// if the function modifies the "existing" elements of the slice,
// the caller will see / observe the changes.
// If the function adds new elements to the slice, that requires changing the slice header
// (the length at a minimum, but may also involve allocating a new backing array),
// which the caller will not see (not without returning the new slice header).

func DfsCaller() {

	// directions := [][]int{
	// 	{-1, 0}, // up
	// 	{0, 1},  // right
	// 	{1, 0},  // down
	// 	{0, -1}, // left
	// }

	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	visited := [][]bool{}
	values := []int{}
	tempVisited := []bool{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			tempVisited = append(tempVisited, false)
		}
		visited = append(visited, tempVisited)
		tempVisited = []bool{}
	}

	dfs(input, 0, 0, visited, &values)
	fmt.Print(values)
}

// Depth first search over 2D array.
func dfs(matrix [][]int, row int, col int, visited [][]bool, values *[]int) {

	if row < 0 || row > len(matrix)-1 || col < 0 || col >= len(matrix[0]) || visited[row][col] {

		return
	}

	*values = append(*values, matrix[row][col])
	visited[row][col] = true

	dfs(matrix, row-1, col, visited, values) // Up
	dfs(matrix, row, col+1, visited, values) // Right
	dfs(matrix, row+1, col, visited, values) // Down
	dfs(matrix, row, col-1, visited, values) // Left
}
