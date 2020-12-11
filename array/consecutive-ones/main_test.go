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
		{in: []int{1, 1, 1, 1}, out: 4},
		{in: []int{0, 0, 0, 0}, out: 0},
		{in: []int{0, 0, 0, 1}, out: 1},
		{in: []int{1, 0, 0}, out: 1},
		{in: []int{1, 1, 1, 0, 0, 1, 1}, out: 3},
	}

	for _, test := range cases {
		res := consecutiveOnes(test.in)
		if res != test.out {
			t.Errorf("consecutiveOnes(%v) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
