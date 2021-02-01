// Mikhail Mir
package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/input"
	"github.com/akaKAIN/algorithms_go/mymath"
	"github.com/akaKAIN/algorithms_go/mysort"
	"log"
	"os"
)

func main() {
	//Task1()
	//Task2()
	Task3()
	Task4()
}

// Реализовать перевод издесятичной в двоичную систему
// счисления с использованием стека.
func Task1() {
	fmt.Println("Введите десятичное число для перевода в двоичную систему")
	userInput, err := input.UserInt(os.Stdin)
	if err != nil {
		log.Println(err)
	}

	arr := mymath.ToBinaryStack(userInput)
	fmt.Printf("%d => ", userInput)
	// Достаем знаяения из стека
	for len(arr) > 0 {
		fmt.Print(arr[len(arr)-1])
		arr = arr[:len(arr)-1]
	}
}

// Добавить в программу «реализация стека на основе односвязного списка» проверку
// на выделение памяти. Если память не выделяется, то выводится соответствующее
// сообщение. Постарайтесь создать ситуацию, когда память не будет выделяться
// (добавлением большого количества данных).
func Task2() {}

// Написать программу, которая определяет, является ли введенная скобочная
// последовательность правильной.
// Примеры правильных скобочных выражений: (), ([])(), {}(), ([{}]),
// неправильных — )(, ())({), (, ])}), ([(]) для скобок [,(,{.
// Например: (2+(2*2)) или [2/{5*(4+7)}]
func Task3() {
	fmt.Println("Введите скобочное выражение")
	expresion := input.UserString(os.Stdin)
	checker := mysort.InitBracketChecker(expresion)
	checker.Main()
}

// Создать функцию, копирующую односвязный список (то есть создает в памяти копию
// односвязного списка, не удаляя первый список).
func Task4() {

}
