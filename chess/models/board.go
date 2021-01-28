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
	Steps []*Step
}

// Инициализация структуры игрового поля.
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
		firstStep := InitStep(basePosition, board.Body)
		board.AddStep(firstStep)

		return board
	} else {
		log.Println("Wrong figure position", basePosition)
		return nil
	}
}

func (b *Board) GetNextStep() *Step {
	lastStep := b.Steps[len(b.Steps)-1]
	if len(lastStep.PossiblePositions) != 0 {
		nextPosition := lastStep.PossiblePositions[len(lastStep.PossiblePositions)-1]
		if isMarked, ok := b.Body[nextPosition]; ok &&
			!isMarked &&
			len(lastStep.PossiblePositions) != 0 {
			return InitStep(nextPosition, b.Body)
		}
	}
	return nil
}

func (b *Board) AddStep(step *Step) {
	b.Steps = append(b.Steps, step)
	b.Body[step.Position] = true
}

func (b *Board) RevertStep() {
	if len(b.Steps) > 1 {
		revertPosition := b.Steps[len(b.Steps)-1].Position
		b.Steps = b.Steps[:len(b.Steps)-1]
		b.Body[revertPosition] = false
		b.Steps[len(b.Steps)-1].DeleteLastPosition()
	} else if len(b.Steps) == 1 {
		if len(b.Steps[0].PossiblePositions) != 0 {
			b.Steps[0].DeleteLastPosition()
		} else {

			log.Fatal("No steps: ", b)
		}
	}
}

func (b *Board) Main() {
	for len(b.Steps) < b.Size.Row*b.Size.Col {
		if nextStep := b.GetNextStep(); nextStep != nil {
			b.AddStep(nextStep)
		} else {
			b.RevertStep()
		}
	}
}
