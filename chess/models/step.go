package models

import "math"

type Step struct {
	Position          Coordinate
	PossiblePositions []Coordinate
}

func (s *Step) CalcPossiblePositions(body map[Coordinate]bool) {
	rowsStepDiff := []int{-2, -1, 1, 2}
	colsStepDiff := []int{-2, -1, 1, 2}
	checkNums := func(num1, num2 int) bool {
		return math.Abs(float64(num1))-math.Abs(float64(num2)) != 0
	}
	for _, rowDiff := range rowsStepDiff {
		for _, colDiff := range colsStepDiff {
			if checkNums(rowDiff, colDiff) {
				newRow := s.Position.Row + rowDiff
				newCol := s.Position.Col + colDiff
				coordinate := Coordinate{Row: newRow, Col: newCol}
				if isMasked, ok := body[coordinate]; ok && !isMasked {
					s.PossiblePositions = append(s.PossiblePositions, coordinate)
				}
			}
		}
	}
}

func (s *Step) DeleteLastPosition() bool {
	if len(s.PossiblePositions) > 0 {
		s.PossiblePositions = s.PossiblePositions[:len(s.PossiblePositions)-1]
		return true
	}
	return false
}

func InitStep(position Coordinate, body map[Coordinate]bool) *Step {
	step := new(Step)
	step.Position = position
	step.CalcPossiblePositions(body)
	return step
}
