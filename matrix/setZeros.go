package matrix

import "fmt"

// Leet code link-> https://leetcode.com/problems/set-matrix-zeroes/
func setZeroes(matrix [][]int) {

	rowLength := len(matrix)
	colLength := len(matrix[0])

	if rowLength == 0 || colLength == 0 {

		return
	}

	zeroPositions := []indices{}
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			if matrix[i][j] == 0 {
				zeroPositions = append(zeroPositions, indices{row: i, col: j})
			}
		}
	}

	for _, zeroPosition := range zeroPositions {

		horizontalDfsRight(matrix, zeroPosition.row, zeroPosition.col)
		horizontalDfsLeft(matrix, zeroPosition.row, zeroPosition.col)

		verticallDfsUp(matrix, zeroPosition.row, zeroPosition.col)
		verticallDfsDown(matrix, zeroPosition.row, zeroPosition.col)
	}
}

func horizontalDfsRight(matrix [][]int, row int, col int) {

	if col < 0 || col >= len(matrix[0]) {

		return
	}

	matrix[row][col] = 0

	// special case for overlapping 0s, we jump one index
	if col+1 < len(matrix[0]) {
		i := 1
		for col+i < len(matrix[0]) && matrix[row][col+i] == 0 {

			i++
		}

		if col+i <= len(matrix[0]) {
			horizontalDfsRight(matrix, row, col+1)
		}
	}

}

func horizontalDfsLeft(matrix [][]int, row int, col int) {

	if col < 0 || col >= len(matrix[0]) {

		return
	}

	matrix[row][col] = 0

	// special case for overlapping 0s, we jump one index
	if col-1 >= 0 {

		j := 1
		for col-j >= 0 && matrix[row][col-j] == 0 {

			j++
		}

		if col-j >= 0 {

			horizontalDfsLeft(matrix, row, col-j)
		}
	}
}

func verticallDfsUp(matrix [][]int, row int, col int) {

	if row < 0 || row >= len(matrix) {

		return
	}

	matrix[row][col] = 0

	// special case for overlapping 0s
	if row+1 < len(matrix) {
		i := 1
		for row+i < len(matrix) && matrix[row+i][col] == 0 {
			i++
		}

		if row+i < len(matrix) {
			verticallDfsUp(matrix, row+i, col)
		}
	}

}

func verticallDfsDown(matrix [][]int, row int, col int) {

	if row < 0 || row >= len(matrix) {

		return
	}

	matrix[row][col] = 0

	// special case for overlapping 0s
	if row-1 >= 0 {
		i := 1
		for row-i >= 0 && matrix[row-i][col] == 0 {
			i++
		}

		if row-i >= 0 {
			verticallDfsDown(matrix, row-i, col)
		}
	}
}

// type indices struct {
// 	row int
// 	col int
// }

func SetZeroCaller() {

	// matrix := [][]int{
	// 	{1, 1, 1},
	// 	{1, 0, 1},
	// 	{1, 1, 1},
	// }
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 0, 7, 8},
		{0, 10, 11, 12},
		{13, 14, 15, 0},
	}
	printAsMatrix(matrix)

	setZeroes(matrix)

	printAsMatrix(matrix)
}

func printAsMatrix(matrix [][]int) {

	fmt.Println("_________________________")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("_________________________")
}
