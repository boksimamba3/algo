package main

import (
	"testing"
)

type TestCase struct {
	in  []int
	out int
}

func TestGetWater(t *testing.T) {

	cases := []TestCase{
		{in: []int{2, 0, 2}, out: 2},
		{in: []int{3, 0, 2, 1, 5}, out: 6},
		{in: []int{3, 0, 3, 0, 2}, out: 5},
		{in: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 3, 1}, out: 8},
	}

	for _, test := range cases {
		res := getWater(test.in)
		if res != test.out {
			t.Errorf("getWater(%v) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
