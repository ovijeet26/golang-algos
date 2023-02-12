package backtracking

// Leet code -> https://leetcode.com/problems/sudoku-solver/
func solveSudoku(board [][]byte) {

	size := len(board)

	present := struct{}{}
	rowsMap := make([]map[int]struct{}, size)
	colsMap := make([]map[int]struct{}, size)
	boxMap := make([]map[int]struct{}, 9)

	for i := 0; i < size; i++ {

		rowsMap[i] = make(map[int]struct{}, 0)
		colsMap[i] = make(map[int]struct{}, 0)
		boxMap[i] = make(map[int]struct{}, 0)
	}

	for row := 0; row < size; row++ {

		for col := 0; col < size; col++ {

			value := board[row][col]

			if value != '.' {
				valueInt := int(value - '0')
				rowsMap[row][valueInt] = present
				colsMap[col][valueInt] = present

				boxId := calculateBoxId(row, col)

				boxMap[boxId][valueInt] = present
			}
		}
	}

	backtrackSolve(board, rowsMap, colsMap, boxMap, 0, 0)
}

func calculateBoxId(row int, col int) int {

	colValue := col / 3
	rowValue := (row / 3) * 3

	return colValue + rowValue
}

func isValid(row int, col int, value int, rowsMap map[int]struct{}, colsMap map[int]struct{}, boxMap map[int]struct{}) bool {

	if _, ok := rowsMap[value]; ok {

		return false
	}

	if _, ok := colsMap[value]; ok {

		return false
	}

	if _, ok := boxMap[value]; ok {

		return false
	}

	return true
}

func backtrackSolve(board [][]byte, rowsMap []map[int]struct{}, colsMap []map[int]struct{}, boxMap []map[int]struct{}, row int, col int) bool {

	// Our solution is complete.
	if row > 8 {

		return true
	}

	// We completed one row.
	// Move to the next row and set column as 0.
	if col > 8 {

		return backtrackSolve(board, rowsMap, colsMap, boxMap, row+1, 0)
	}

	currentBoardValue := board[row][col]
	// If the board cell is already pre-filled, we skip that column.
	if currentBoardValue != '.' {

		return backtrackSolve(board, rowsMap, colsMap, boxMap, row, col+1)

	} else {

		for i := 1; i <= 9; i++ {

			boxId := calculateBoxId(row, col)

			if isValid(row, col, i, rowsMap[row], colsMap[col], boxMap[boxId]) {

				board[row][col] = byte(i + '0')

				rowsMap[row][i] = struct{}{}
				colsMap[col][i] = struct{}{}
				boxMap[boxId][i] = struct{}{}

				if backtrackSolve(board, rowsMap, colsMap, boxMap, row, col+1) == true {

					return true
				}

				board[row][col] = '.'
				delete(rowsMap[row], i)
				delete(colsMap[col], i)
				delete(boxMap[boxId], i)

			}
		}
		return false
	}
}

func SodokuCaller() {

	grid := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(grid)
}

func solveSudokuNew(board [][]byte) {
	// divide board into 9 subboard from 0 to 8
	// try to fill each subboard
	// initialize rows, cols, subs 2D arrays to keep track which number already exist
	// for example rows[3][6] is true means the 4th row(index 3) already have 7(index 6) in it
	rows, cols, subs := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i], cols[i], subs[i] = make([]bool, 9), make([]bool, 9), make([]bool, 9)
	}

	// search existing board, fill all rows, cols and subs
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				rows[i][num-1] = true
				cols[j][num-1] = true
				sub := (i/3)*3 + j/3
				subs[sub][num-1] = true
			}
		}
	}

	var backtracking func(int, int) bool
	backtracking = func(x, y int) bool {
		// when row index > 8 means, we already filled all first 9 row, then done
		if x > 8 {
			return true
		}

		// we go from left to righ, top to bottom, and next location we will be visiting is nextX,nextY
		nextX := x + (y+1)/9
		nextY := (y + 1) % 9
		if board[x][y] == '.' {
			// when current position is empty, we try to fill it from 1 to 9
			for num := 1; num <= 9; num++ {
				sub := (x/3)*3 + y/3
				// if the number we are trying is not valid, move to next
				if rows[x][num-1] || cols[y][num-1] || subs[sub][num-1] {
					continue
				}
				//fill the number
				board[x][y] = byte('0' + num)
				rows[x][num-1] = true
				cols[y][num-1] = true
				subs[sub][num-1] = true

				// stop further process if backtracking returns true
				if backtracking(nextX, nextY) {
					return true
				}
				//remove the number, so we can try next one
				board[x][y] = '.'
				rows[x][num-1] = false
				cols[y][num-1] = false
				subs[sub][num-1] = false
			}
			return false

		} else {
			return backtracking(nextX, nextY)
		}
	}
	backtracking(0, 0)
}
