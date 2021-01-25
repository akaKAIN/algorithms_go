package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinary(t *testing.T) {
	tests := []struct {
		inArr []byte
		inVal byte
		out   int
		name  string
	}{
		{inArr: []byte("123456"), inVal: '4', out: 3, name: "One"},
		{inArr: []byte("123456"), inVal: '2', out: 1, name: "Two"},
		{inArr: []byte("0123456789"), inVal: '4', out: 4, name: "Three"},
		{inArr: []byte("0123456789"), inVal: '6', out: 6, name: "Four"},
		{inArr: []byte("0123456789"), inVal: '0', out: 0, name: "Five"},
		{inArr: []byte("0123456789"), inVal: '9', out: 9, name: "Six"},
	}

	for _, tc := range tests {
		result := Binary(tc.inArr, tc.inVal)
		assert.Equal(t, tc.out, result, tc.name)
	}
}
