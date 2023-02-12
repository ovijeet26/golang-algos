package matrix

import "fmt"

// Leet code link -> https://leetcode.com/problems/word-search/
func exist(board [][]byte, word string) bool {

	rowSize := len(board)
	colSize := len(board[0])
	indexArray := make([]index, 0)
	visited := make([][]bool, rowSize)

	// Linear scan the board to find starting index occurances
	for i := 0; i < rowSize; i++ {

		for j := 0; j < colSize; j++ {

			if board[i][j] == word[0] {

				indexArray = append(indexArray, index{row: i, col: j})
			}
		}

		// Make the visited matrix
		visitedRow := make([]bool, colSize)
		visited[i] = visitedRow
	}

	// Check if no staring index is present, then skip compute
	if len(indexArray) == 0 {

		return false
	}

	// Iterate over list of starting points to run dfsSearch
	for _, index := range indexArray {

		currentRow := index.row
		currentCol := index.col

		isFound := dfsSearch(board, visited, word, currentRow, currentCol, 0)

		if isFound {

			return true
		}
	}

	return false
}

type index struct {
	row int
	col int
}

func dfsSearch(board [][]byte, visited [][]bool, word string, row int, col int, index int) bool {

	// Check boundary conditions
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || visited[row][col] {

		return false
	}

	// If the word is not satisfied in this path
	if word[index] != board[row][col] {

		return false
	}

	// Marking teh current cell to be visited
	visited[row][col] = true

	// If our trersal found the word. Success condition.
	if index == len(word)-1 {

		return true
	}

	topPass := dfsSearch(board, visited, word, row+1, col, index+1)
	leftPass := dfsSearch(board, visited, word, row, col-1, index+1)
	rightPass := dfsSearch(board, visited, word, row, col+1, index+1)
	downPass := dfsSearch(board, visited, word, row-1, col, index+1)

	// Backtrack to retry
	visited[row][col] = false

	return (topPass || leftPass || downPass || rightPass)
}

func WordSearchCaller() {

	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }

	// word := "ABCCED"

	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	word := "AAB"

	fmt.Println(exist(board, word))
}

// search -> abcd

// [a b c x]
// [v s d s]
// [b a w x]

// Assumptions ->
// 1) Will there be repetaion? Yes
// 2) Only characters
// 3) search string can be backwards -> abcd \ dbca

// Approach -> get position of start index
