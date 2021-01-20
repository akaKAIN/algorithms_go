// student: Mikhail Mir
package main

import (
	"bufio"
	"fmt"
	"github.com/akaKAIN/algorithms_go/input"
	"github.com/akaKAIN/algorithms_go/mymath"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Выберите номер пункта для ввода данных:")
	fmt.Print("1. Ввод с клавиатуры\n2. С файла\n")
	var startNum, limit int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "1":
		fmt.Println("Введите стартовое число: ")
		num, err := input.UserInt(os.Stdin)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		startNum = num
		fmt.Println("Введите пороговое число для расчетов: ")
		limit, err = input.UserInt(os.Stdin)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Println(
			"Найдено решений:",
			mymath.WayCounterRecursive(startNum, limit),
		)
	case "2":
		fmt.Println("Введите имя файла для чтения: ")
		fileName := input.UserString(os.Stdin)
		fileText, err := input.UserFile(fileName)
		if err != nil {
			log.Println("File opening error: ", err)
		}
		words := strings.Split(fileText, " ")
		startNum, err = strconv.Atoi(words[0])
		if err != nil {
			log.Println(err)
		}
		limit, err = strconv.Atoi(words[1])
		if err != nil {
			log.Println(err)
		}
		fmt.Println(
			"Найдено решений:",
			mymath.WayCounterRecursive(startNum, limit),
		)
	}
	fmt.Println(startNum, limit)
}
