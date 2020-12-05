package main

import "testing"

type TestCase struct {
	in  int
	out int
}

func TestPositiveNumber(t *testing.T) {

	cases := []TestCase{
		{10, 2},
		{9567, 4},
		{342, 3},
		{8348394, 7},
	}

	for _, test := range cases {
		res := countDigits(test.in)
		if test.out != res {
			t.Errorf("countDigits(%d) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}

func TestNegativeNumber(t *testing.T) {

	cases := []TestCase{
		{-10, 0},
		{-9567, 0},
		{-342, 0},
		{-8348394, 0},
	}

	for _, test := range cases {
		res := countDigits(test.in)
		if test.out != res {
			t.Errorf("countDigits(%d) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
