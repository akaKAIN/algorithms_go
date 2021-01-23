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

func UserInt(in io.Reader) (int, error)  {
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
