package cripto

import "fmt"

func SimpleCripto(in string) string {
	var hash string
	inByteArr := []byte(in)
	for ind := range inByteArr {
		hash += fmt.Sprintf(
			"%dx%v",
			inByteArr[ind]+byte(ind+1*len(inByteArr)),
			string(inByteArr[ind]+byte(len(inByteArr))),
		)
	}
	return hash
}
