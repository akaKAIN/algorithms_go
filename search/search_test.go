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

func TestStepCalc(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{in: 3, out: 1},
		{in: 1, out: 1},
		{in: 4, out: 2},
		{in: 10, out: 5},
		{in: 7, out: 3},
	}
	for _, tc := range tests {
		result := stepCalc(tc.in)
		assert.Equal(t, tc.out, result)
	}
}
