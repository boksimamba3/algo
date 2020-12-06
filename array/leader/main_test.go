package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	in  []int
	out []int
}

func TestGetLeaders(t *testing.T) {

	cases := []TestCase{
		{in: []int{7, 10, 4, 3, 6, 5, 2}, out: []int{2, 5, 6, 10}},
		{in: []int{10, 20, 30}, out: []int{30}},
		{in: []int{30, 20, 10}, out: []int{10, 20, 30}},
	}

	for _, test := range cases {
		res := getLeaders(test.in)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("getLeaders(%v) failed, expected %v got %v", test.in, test.out, res)
		}
	}
}
