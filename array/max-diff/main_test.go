package main

import (
	"testing"
)

type TestCase struct {
	in  []int
	out int
}

func TestMaxDiff(t *testing.T) {

	cases := []TestCase{
		{in: []int{2, 3, 10, 6, 4, 8, 1}, out: 8},
		{in: []int{7, 9, 5, 6, 3, 2}, out: 2},
		{in: []int{10, 20, 30}, out: 20},
		{in: []int{30, 10, 8, 2}, out: -2},
		{in: []int{2, 3, 10, 6, 4, 17}, out: 15},
	}

	for _, test := range cases {
		res := maxDiff(test.in)
		if res != test.out {
			t.Errorf("maxDiff(%v) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
