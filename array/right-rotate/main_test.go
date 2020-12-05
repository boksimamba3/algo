package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	in  []int
	d   int // number of right rotations
	out []int
}

func TestReverse(t *testing.T) {

	cases := []TestCase{
		{in: []int{1, 2, 3, 4, 5}, d: 1, out: []int{5, 1, 2, 3, 4}},
		{in: []int{1, 2, 3, 4, 5}, d: 2, out: []int{4, 5, 1, 2, 3}},
		{in: []int{1, 2, 3, 4, 5}, d: 3, out: []int{3, 4, 5, 1, 2}},
		{in: []int{1, 2, 3, 4, 5}, d: 4, out: []int{2, 3, 4, 5, 1}},
	}

	for _, test := range cases {
		original := make([]int, len(test.in))
		copy(original, test.in)
		rightRotate(test.in, test.d)
		if !reflect.DeepEqual(test.in, test.out) {
			t.Errorf("rightRotate(%v, %d) failed, expected %v got %v", original, test.d, test.out, test.in)
		}
	}
}
