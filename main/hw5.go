// Mikhail Mir
package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/input"
	"github.com/akaKAIN/algorithms_go/mymath"
	"log"
	"os"
)

func main() {
	Task1()
	Task2()
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

func Task2() {}
func Task3() {}
func Task4() {}
