package models

import (
	"log"
)

type Cell struct {
	Coordinate Coordinate
	isMarked   bool
}

type Board struct {
	Body   []Cell
	Figure *Knight
}

// Инициализация структуры доски.
func InitBoard(row, col int, basePosition [2]int) *Board {
	if row >= basePosition[0] || col >= basePosition[1] {
		board := new(Board)
		for i := 0; i <= row; i++ {
			for j := 0; j <= col; j++ {
				coordinate := Coordinate{Row: i, Col: j}
				board.Body = append(board.Body, Cell{Coordinate: coordinate, isMarked: false})
			}
		}
		board.Figure = InitKnight(Coordinate{Row: basePosition[0], Col: basePosition[1]})
		return board
	} else {
		log.Println("Wrong figure position", basePosition)
		return nil
	}
}
