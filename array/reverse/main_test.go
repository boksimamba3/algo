package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	in  []int
	out []int
}

func TestReverse(t *testing.T) {

	testData := []TestData{
		{in: []int{20, 10, 20, 8, 12}, out: []int{12, 8, 20, 10, 20}},
		{in: []int{10, 10, 10}, out: []int{10, 10, 10}},
		{in: []int{1, 2, 3, 4, 5}, out: []int{5, 4, 3, 2, 1}},
	}

	for _, test := range testData {
		original := make([]int, len(test.in))
		copy(original, test.in)
		reverse(test.in)
		if !reflect.DeepEqual(test.in, test.out) {
			t.Errorf("reverse(%v) failed, expected %v got %v", original, test.out, test.in)
		}
	}
}
