package mymath

func ToBinary(num int) int {
	if num == 2 {
		return 0
	} else if num == 1 {
		return 1
	}
	return num % 2 + 10 * ToBinary(num/2)
}
