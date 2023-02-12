package matrix

import "fmt"

// Leet code link -> https://leetcode.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {

	res := make([]int, 0)

	left := 0
	right := len(matrix[0])

	top := 0
	bottom := len(matrix)

	for top < bottom && left < right {

		for i := left; i < right; i++ {

			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i < bottom; i++ {

			res = append(res, matrix[i][right-1])
		}

		right--

		if !(top < bottom && left < right) {

			break
		}

		for i := right - 1; i >= left; i-- {

			res = append(res, matrix[bottom-1][i])
		}

		bottom--
		for i := bottom - 1; i > top-1; i-- {

			res = append(res, matrix[i][left])
		}

		left++
	}

	return res
}

func SpiralMatrixCaller() {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	res := spiralOrder(matrix)

	fmt.Println(res)
}
