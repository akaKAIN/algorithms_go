package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/cripto"
	"github.com/akaKAIN/algorithms_go/input"
	"os"
)

func main() {
	HomeWork6Task1()
}

// Реализовать простейшую хеш-функцию. На вход функции подается строка, на выходе сумма кодов символов.
func HomeWork6Task1() {
	fmt.Println("Введите строку для хэширования:")
	UserInput := input.UserString(os.Stdin)
	hash := cripto.SimpleCripto(UserInput)
	fmt.Println(hash)

}
