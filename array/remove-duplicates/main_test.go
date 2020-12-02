package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	in  []int
	out []int
}

func TestRemoveDuplicates(t *testing.T) {

	testData := []TestData{
		{in: []int{10, 20, 20, 30, 30, 30, 30, 40, 40, 50, 60}, out: []int{10, 20, 30, 40, 50, 60}},
		{in: []int{10, 10, 10}, out: []int{10}},
		{in: []int{10, 40, 40}, out: []int{10, 40}},
		{in: []int{}, out: []int{}},
	}

	for _, test := range testData {
		res := removeDuplicates(test.in)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("reverse(%v) failed, expected %v got %v", test.in, test.out, res)
		}
	}
}
