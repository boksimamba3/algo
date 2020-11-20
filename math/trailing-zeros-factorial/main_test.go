package main

import "testing"

type TestData struct {
	in  int
	out int
}

func TestCountFactorialTrailingZeros(t *testing.T) {
	testData := []TestData{
		{5, 1},
		{10, 2},
		{251, 62},
	}

	for _, test := range testData {
		res := countZeros(test.in)
		if test.out != res {
			t.Errorf("countZeros(%d) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
