package search

func stepCalc(in int) int {
	if in == 1 {
		return 1
	}
	return in / 2
}

func Binary(sorterArr []byte, value byte) int {
	var (
		idx = len(sorterArr) / 2
		step = stepCalc(idx)
	)

	for {
		step = stepCalc(step)
		if sorterArr[idx] == value {
			return idx
		} else if sorterArr[idx] > value {
			idx -= step
		} else if sorterArr[idx] < value {
			idx += step
		}
		if sorterArr[idx] == value {
			return idx
		}
		if idx == 0 {
			return -1
		}
	}
}
