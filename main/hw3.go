// Mikhail Mir
package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/input"
	"os"
)

func main() {
	fmt.Println("Ввеберите пункт:")
	fmt.Println("1. Сортировка переданной строки")
	fmt.Println("2. Бинарный поиск")
	userInput := input.UserString(os.Stdin)
	fmt.Println(userInput)
}
