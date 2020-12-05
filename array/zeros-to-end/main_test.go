package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	in  []int
	out []int
}

func TestZerosToEnd(t *testing.T) {

	cases := []TestCase{
		{in: []int{20, 10, 20, 8, 12}, out: []int{20, 10, 20, 8, 12}},
		{in: []int{10, 10, 0, 0, 10}, out: []int{10, 10, 10, 0, 0}},
		{in: []int{1, 0, 0, 2, 3}, out: []int{1, 2, 3, 0, 0}},
		{in: []int{8, 5, 0, 10, 0, 20}, out: []int{8, 5, 10, 20, 0, 0}},
		{in: []int{0, 0, 0, 10, 0}, out: []int{10, 0, 0, 0, 0}},
	}

	for _, test := range cases {
		original := make([]int, len(test.in))
		copy(original, test.in)
		zerosToEnd(test.in)
		if !reflect.DeepEqual(test.in, test.out) {
			t.Errorf("zerosToEnd(%v) failed, expected %v got %v", original, test.out, test.in)
		}
	}
}
