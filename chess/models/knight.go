package models

// Положение по вертикали (Row) и горизонтали(Col)
type Coordinate struct {
	Row int
	Col int
}

// Пложение фигуры на доске. Хранит информацию о направлении предыдущего и последующего ходов.
// Под "направлением" следует понимать порядковый номер одного из 8-ми возможных направлений.
type Step struct {
	PreviousDirection int
	NextDirection     int
	Position          Coordinate
}

// Положение фигуры на доск (CurrentPosition) и массив с ее вероятными позициями на ней (PossibleDirections).
type Knight struct {
	CurrentPosition    Coordinate
	PossibleDirections []Coordinate
}

// Метод изменения положения фигуры и расчета дальнейших возможных ходов.
func (k *Knight) MoveTo(position Coordinate) {
	k.CurrentPosition = position
	k.calcPossibleDirections()
}

// Подстановка ВСЕХ возможных ходов.
func (k *Knight) calcPossibleDirections() {
	k.PossibleDirections = []Coordinate{
		{Row: k.CurrentPosition.Row + 2, Col: k.CurrentPosition.Col + 1},
		{Row: k.CurrentPosition.Row + 2, Col: k.CurrentPosition.Col - 1},
		{Row: k.CurrentPosition.Row + 1, Col: k.CurrentPosition.Col + 2},
		{Row: k.CurrentPosition.Row + 1, Col: k.CurrentPosition.Col - 2},
		{Row: k.CurrentPosition.Row - 2, Col: k.CurrentPosition.Col + 1},
		{Row: k.CurrentPosition.Row - 2, Col: k.CurrentPosition.Col - 1},
		{Row: k.CurrentPosition.Row - 1, Col: k.CurrentPosition.Col + 2},
		{Row: k.CurrentPosition.Row - 1, Col: k.CurrentPosition.Col - 2},
	}
}

func InitKnight(coordinate Coordinate) *Knight {
	knight := new(Knight)
	knight.CurrentPosition = coordinate
	knight.calcPossibleDirections()
	return knight
}
