package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/chess/models"
)

func main() {
	board := models.InitBoard(8, 8, models.Coordinate{Row: 1, Col: 3})
	board.Main()
	fmt.Println("Board:")
	for key, val := range board.Body {
		fmt.Println(key, val)
	}
	fmt.Printf("steps: %q\n", board.Steps)
}
