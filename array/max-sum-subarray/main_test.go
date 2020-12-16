package main

import (
	"testing"
)

type TestCase struct {
	in  []int
	out int
}

func TestMaxSumSubarray(t *testing.T) {

	cases := []TestCase{
		{in: []int{-5, 1, -2, 3, -1, 2, -2}, out: 4},
		{in: []int{2, 3, -8, 7, -1, 2, 3, -1, 4}, out: 14},
		{in: []int{-2, -3, 4, -1, -2, 1, 5, -3}, out: 7},
	}

	for _, test := range cases {
		res := maxSumSubarray(test.in)
		if res != test.out {
			t.Errorf("maxSumSubarray(%v) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
