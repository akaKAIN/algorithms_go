package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/chess/models"
)

func main() {
	board := models.InitBoard(4, 3, models.Coordinate{Row: 1, Col: 1})
	board.AddStep(board.Steps[0].PossiblePositions[len(board.Steps[0].PossiblePositions)-1])
	fmt.Println(board)
}
