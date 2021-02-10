package mysort

import (
	"bytes"
	"log"
)

type BracketChecker struct {
	possibleInputBracket []byte
	Expresion            []byte
	InputStack           []byte
	ClosedBracket        map[byte]byte
}

func InitBracketChecker(startExpresion string) *BracketChecker {
	bracket := new(BracketChecker)
	bracket.Expresion = []byte(startExpresion)
	bracket.possibleInputBracket = []byte{'[', '(', '{'}
	bracket.ClosedBracket = make(map[byte]byte)
	bracket.ClosedBracket[']'] = '['
	bracket.ClosedBracket[')'] = '('
	bracket.ClosedBracket['}'] = '{'
	return bracket
}

func (b *BracketChecker) Add(bracket byte) {
	b.InputStack = append(b.InputStack, bracket)
}

func (b *BracketChecker) isCorrectClosing(bracket byte) bool {
	if len(b.InputStack) > 0 {
		lastBracket := b.InputStack[len(b.InputStack)-1]
		if b.ClosedBracket[bracket] == lastBracket {
			return true
		}
		return false
	}
	return false
}

func (b *BracketChecker) clearLastBracket() {
	b.InputStack = b.InputStack[:len(b.InputStack)-1]
}

func (b *BracketChecker) Main() {
	isExpresionValid := true
	for _, bracket := range b.Expresion {
		if bytes.ContainsAny(b.possibleInputBracket, string(bracket)) {
			b.Add(bracket)
		} else if b.isCorrectClosing(bracket) {
			b.clearLastBracket()
		} else {
			isExpresionValid = false
			break
		}

	}
	if isExpresionValid {
		log.Println("Выражение верное.")
	} else {
		log.Println("Выражение НЕ верное.")
	}
}

func (b *BracketChecker) IsValidBracket(bracket byte) bool {
	if _, ok := b.ClosedBracket[bracket]; !ok {
		if correctClosing := b.isCorrectClosing(bracket); correctClosing {
			return true
		} else {
			return false
		}
	}
	return true

}

type Queue struct {
	Body []byte
}

func (q *Queue) Pop() byte {
	limit := len(q.Body) - 1
	last := q.Body[limit]
	q.Body = q.Body[:limit]
	return last
}

func (q *Queue) Put(data byte) {
	body := make([]byte, len(q.Body), len(q.Body)+1)
	copy(body, q.Body)
	q.Body = append([]byte{data}, body...)
}
