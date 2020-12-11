package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	in  []int
	out [][]int
}

func TestFrequenciesCounter(t *testing.T) {

	cases := []TestCase{
		{
			in:  []int{1, 1, 1, 1, 2, 2, 3, 3},
			out: [][]int{{1, 4}, {2, 2}, {3, 2}},
		},
		{
			in:  []int{1, 1, 1, 1},
			out: [][]int{{1, 4}},
		},
	}

	for _, test := range cases {
		res := frequenciesCounter(test.in)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("frequenciesCounter(%v) failed, expected %v got %v", test.in, test.out, res)
		}
	}
}
