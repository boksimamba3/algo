package main

import "testing"

type TestCase struct {
	n  int // number
	dp int // digit position
	d  int // digit
}

func TestGetNthDigit(t *testing.T) {

	cases := []TestCase{
		{10, 1, 0},
		{10, 2, 1},
		{9567, 4, 9},
		{9567, 2, 6},
		{8348394, 5, 4},
		{8348394, 3, 3},
	}

	for _, test := range cases {
		res := getDigit(test.n, test.dp)
		if test.d != res {
			t.Errorf("getDigit(%d, %d) failed, expected %d got %d", test.n, test.dp, test.d, res)
		}
	}
}
