package mysort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	tests := []struct {
		inArr  []byte
		outArr []byte
		name   string
	}{
		{inArr: []byte("qaz"), outArr: []byte("aqz"), name: "word: qaz"},
		{inArr: []byte("gdcbfea"), outArr: []byte("abcdefg"), name: "word: gdcbfea"},
		{inArr: []byte("54321"), outArr: []byte("12345"), name: "54321"},
		{inArr: []byte("744261"), outArr: []byte("124467"), name: "744261"},
		{inArr: []byte("12345"), outArr: []byte("12345"), name: "12345"},
	}

	for _, tc := range tests {
		resultArr := Bubble(tc.inArr)
		assert.Equal(t, tc.outArr, resultArr, tc.name)
	}
}

func TestShaker(t *testing.T) {
	tests := []struct {
		inArr  []byte
		outArr []byte
		name   string
	}{
		{inArr: []byte("qaz"), outArr: []byte("aqz"), name: "word: qaz"},
		{inArr: []byte("gdcbfea"), outArr: []byte("abcdefg"), name: "word: gdcbfea"},
		{inArr: []byte("54321"), outArr: []byte("12345"), name: "54321"},
		{inArr: []byte("744261"), outArr: []byte("124467"), name: "744261"},
		{inArr: []byte("744267"), outArr: []byte("244677"), name: "744267"},
		{inArr: []byte("12345"), outArr: []byte("12345"), name: "12345"},
	}

	for _, tc := range tests {
		resultArr := Shaker(tc.inArr)
		assert.Equal(t, tc.outArr, resultArr, tc.name)
	}
}
