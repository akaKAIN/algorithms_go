package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/cripto"
	"github.com/akaKAIN/algorithms_go/input"
	"github.com/akaKAIN/algorithms_go/tree"
	"os"
)

func main() {
	//HomeWork6Task1()
	HomeWork6Task2()
}

// Реализовать простейшую хеш-функцию. На вход функции подается строка, на выходе сумма кодов символов.
func HomeWork6Task1() {
	fmt.Println("Введите строку для хэширования:")
	UserInput := input.UserString(os.Stdin)
	hash := cripto.SimpleCripto(UserInput)
	fmt.Println(hash)
}

// Переписать программу, реализующую двоичное дерево поиска.
func HomeWork6Task2() {
	Tree := tree.InitNode([]byte("First"))
	Tree.AddData([]byte("Hi"))
	Tree.AddData([]byte("As"))
	Tree.AddData([]byte("Ab"))
	fmt.Println(Tree)
}
