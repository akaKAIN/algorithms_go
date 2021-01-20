package mymath

func ToBinary(num int) int {
	if num == 0 {
		return 0
	}
	return num%2 + 10*ToBinary(num/2)
}

func PowRecursive(num, pow int) int {
	if pow == 0 {
		return 1
	}
	return num * PowRecursive(num, pow-1)
}

func Pow(num, pow int) int {
	var result = 1
	for pow > 0 {
		result *= num
		pow--
	}
	return result
}

func PowRecursiveEven(num, pow int) int {
	if pow == 0 {
		return 1
	}

	if pow%2 == 0 {
		return PowRecursiveEven(num, pow/2) * PowRecursiveEven(num, pow/2)
	} else {
		return num * PowRecursiveEven(num, pow-1)
	}
}

func WayCounterRecursive(start, limit int) int {
	if start == limit || start*2 > limit {
		return 1
	}
	if start*2 <= limit {
		return WayCounterRecursive(start+1, limit) + WayCounterRecursive(start*2, limit)
	}
	return 1
}
