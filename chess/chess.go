package chess

import "fmt"

func Horse(rowNum, colNum int) {
	board := generateMatrix(rowNum, colNum)
	markStep(board, 1, 1, 1)
	fmt.Println(board)
}

func generateMatrix(row, col int) [][]int {
	matrix := make([][]int, row, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col, col)
	}
	return matrix
}

func markStep(matrix [][]int, row, col, marker int) bool {
	if matrix[row][col] == marker {
		return false
	} else {
		matrix[row][col] = marker
		return true
	}
}
