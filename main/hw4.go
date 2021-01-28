package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/chess/models"
)

func main() {
	board := models.InitBoard(4, 4, [2]int{1, 1})
	fmt.Println(board)
}
