package main

import (
	"testing"
)

type TestCase struct {
	in  []int
	out int
}

func TestMaxProfit(t *testing.T) {

	cases := []TestCase{
		{in: []int{1, 2, 3, 4, 5}, out: 4},
		{in: []int{5, 4, 3, 2, 1}, out: 0},
		{in: []int{1, 5, 3, 1, 2, 8}, out: 11},
		{in: []int{1, 5, 3, 8, 12}, out: 13},
	}

	for _, test := range cases {
		res := maxProfit(test.in)
		if res != test.out {
			t.Errorf("maxProfit(%v) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
