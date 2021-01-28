package models

import (
	"log"
)

// Положение по вертикали (Row) и горизонтали(Col)
type Coordinate struct {
	Row int
	Col int
}

type SizeBoard = Coordinate

type Board struct {
	Body  map[Coordinate]bool
	Size  SizeBoard
	Steps []Step
}

// Инициализация структуры доски.
func InitBoard(row, col int, basePosition Coordinate) *Board {
	if row >= basePosition.Row || col >= basePosition.Col {
		board := new(Board)
		board.Body = make(map[Coordinate]bool)
		board.Size = SizeBoard{Row: row, Col: col}
		for i := 1; i <= row; i++ {
			for j := 1; j <= col; j++ {
				coordinate := Coordinate{Row: i, Col: j}
				board.Body[coordinate] = false
			}
		}
		wasAdded := board.AddStep(basePosition)
		if !wasAdded {
			log.Println("Some kind of shit is happened in first step")
		}
		return board
	} else {
		log.Println("Wrong figure position", basePosition)
		return nil
	}
}

func (b *Board) GetNextStep() *Step {
	lastStep := b.Steps[len(b.Steps)-1]
	nextPosition := lastStep.PossiblePositions[len(lastStep.PossiblePositions)-1]
	if isMarked, ok := b.Body[nextPosition]; ok &&
		!isMarked &&
		len(lastStep.PossiblePositions) != 0 {
		newStep := new(Step)
		newStep.Position = nextPosition
		newStep.CalcPossiblePositions(b.Body)
		return newStep
	}
	return nil
}

func (b *Board) AddStep(step Step)  {
	b.Steps = append(b.Steps, step)
}

func (b *Board) Main() {
	for len(b.Steps) < b.Size.Row*b.Size.Col {
		if nextStep := b.GetNextStep(); nextStep != nil {
			b.AddStep(*nextStep)
		} else  {

		}
	}
}
