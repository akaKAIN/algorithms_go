package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/cripto"
	"github.com/akaKAIN/algorithms_go/input"
	"github.com/akaKAIN/algorithms_go/tree"
	"log"
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
	func(){
		defer func() {
			if r := recover(); r != nil {
				log.Println("No match")
			}
		}()
	}()
	Tree := tree.InitNode([]byte("D"))
	Tree.AddData([]byte("B"))
	Tree.AddData([]byte("A"))
	Tree.AddData([]byte("F"))
	Tree.AddData([]byte("E"))
	Tree.AddData([]byte("G"))
	fmt.Println(Tree)
	result := tree.Search(Tree, []byte("G"))
	log.Println("Road to target is: ", result)
}
