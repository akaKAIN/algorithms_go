package mysort

// Пузырьковая сортировка
func Bubble(arr []byte) []byte {
	for {
		isSorted := true
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				isSorted = false
			}
		}
		if isSorted {
			return arr
		}
	}
}

// Шейкерная сортировка
func Shaker(arr []byte) []byte {
	var (
		minLimit = 1
		maxLimit = len(arr)
	)
	for {
		isSorted := true
		for i := minLimit; i < maxLimit; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				isSorted = false
			}
		}
		maxLimit--
		for i := maxLimit; i >= minLimit; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				isSorted = false
			}
		}
		minLimit++
		if isSorted {
			return arr
		}
	}
}
