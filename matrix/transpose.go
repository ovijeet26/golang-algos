package matrix

// Leet code link -> https://leetcode.com/problems/transpose-matrix/submissions/
func transpose(matrix [][]int) [][]int {

	rowLength := len(matrix)
	colLength := len(matrix[0])

	if rowLength == 0 || colLength == 0 {

		return matrix
	}

	// Transpose matrix should be of row length of original column
	// and row length should be of original column length.
	transpose := make([][]int, colLength)

	// Instantiate all the rows, but with reversed sizes ( row/col )
	for i := 0; i < colLength; i++ {

		transpose[i] = make([]int, rowLength)
	}

	// Simply make row the column, and vice versa.
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			transpose[j][i] = matrix[i][j]
		}
	}

	return transpose
}
