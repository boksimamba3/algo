package main

import "testing"

type TestCase struct {
	in  int
	out int
}

func TestCountFactorialTrailingZeros(t *testing.T) {
	cases := []TestCase{
		{5, 1},
		{10, 2},
		{251, 62},
	}

	for _, test := range cases {
		res := countZeros(test.in)
		if test.out != res {
			t.Errorf("countZeros(%d) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
