package search

import "fmt"

func Binary(sorterArr []byte, value byte) int {
	var idx = len(sorterArr)
	for {
		if idx % 2 == 0 {
			idx /= 2
		} else {
			idx = (idx-1) / 2
		}
		if sorterArr[idx] == value {
			return idx
		} else if sorterArr[idx] < value {
			sorterArr = sorterArr[idx:]
		} else if sorterArr[idx] > value {
			sorterArr = sorterArr[:idx]

		}
		if idx == 0 {
			return -1
		}
		fmt.Println(idx, string(sorterArr))
	}
}
