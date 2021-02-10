package mymath

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToBinary(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{in: 0, out: 0},
		{in: 1, out: 1},
		{in: 2, out: 10},
		{in: 3, out: 11},
		{in: 4, out: 100},
		{in: 5, out: 101},
	}

	for _, tc := range tests {
		result := ToBinary(tc.in)
		assert.Equal(t, tc.out, result, "Must be equal")
	}
}

func TestPowRecursive(t *testing.T) {
	tests := []struct {
		num int
		pow int
		out int
	}{
		{num: 2, pow: 1, out: 2},
		{num: 2, pow: 2, out: 4},
		{num: 2, pow: 3, out: 8},
		{num: 3, pow: 3, out: 27},
	}
	for _, tc := range tests {
		result := PowRecursive(tc.num, tc.pow)
		assert.Equal(t, tc.out, result, "Must be equal")
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		num int
		pow int
		out int
	}{
		{num: 2, pow: 1, out: 2},
		{num: 2, pow: 2, out: 4},
		{num: 2, pow: 3, out: 8},
		{num: 3, pow: 3, out: 27},
	}
	for _, tc := range tests {
		result := Pow(tc.num, tc.pow)
		assert.Equal(t, tc.out, result, "Must be equal")
	}
}

func TestPowRecursiveEven(t *testing.T) {
	tests := []struct {
		num int
		pow int
		out int
	}{
		{num: 2, pow: 1, out: 2},
		{num: 2, pow: 2, out: 4},
		{num: 2, pow: 3, out: 8},
		{num: 2, pow: 4, out: 16},
		{num: 3, pow: 3, out: 27},
	}
	for _, tc := range tests {
		result := PowRecursiveEven(tc.num, tc.pow)
		assert.Equal(t, tc.out, result, "Must be equal")
	}
}

func TestWayCounter(t *testing.T) {
	tests := []struct {
		start int
		limit int
		out   int
	}{
		{start: 1, limit: 2, out: 2},
		{start: 2, limit: 5, out: 2},
		{start: 3, limit: 10, out: 4},
		{start: 2, limit: 7, out: 3},
	}
	for _, tc := range tests {
		result := WayCounterRecursive(tc.start, tc.limit)
		assert.Equal(t, tc.out, result, "Must be equal")
	}
}

func TestToBinaryStack(t *testing.T) {
	tests := []struct {
		in  int
		out []int
	}{
		{in: 0, out: []int{0}},
		{in: 1, out: []int{1}},
		{in: 2, out: []int{0, 1}},
		{in: 3, out: []int{1, 1}},
		{in: 4, out: []int{0, 0, 1}},
		{in: 5, out: []int{1, 0, 1}},
	}

	for _, tc := range tests {
		result := ToBinaryStack(tc.in)
		assert.Equal(t, tc.out, result, fmt.Sprintf("For %d", tc.in))
	}
}
