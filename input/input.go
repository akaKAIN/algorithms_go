package input

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func UserString(in io.Reader) string {
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	return scanner.Text()
}

func UserInt(in io.Reader) (int, error) {
	strNum := UserString(in)
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func UserArray(in io.Reader, sepr string) []string {
	return strings.Split(UserString(in), sepr)
}

func UserFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return UserString(file), nil
}

func UserFileWithMatrix(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	clearMatrix := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textArr := strings.Split(scanner.Text(), " ")
		intArr := make([]int, len(textArr))
		for i, s := range textArr {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			intArr[i] = num
		}
		clearMatrix = append(clearMatrix, intArr)
	}
	return clearMatrix, nil
}
