package main

import "testing"

type TestData struct {
	in  []int
	out int
}

func TestSecondLargest(t *testing.T) {

	testData := []TestData{
		{in: []int{10}, out: -1},
		{in: []int{10, 5, 8, 20}, out: 0},
		{in: []int{20, 10, 20, 8, 12}, out: 4},
		{in: []int{10, 10, 10}, out: -1},
	}

	for _, test := range testData {
		res := secondLargest(test.in)
		if test.out != res {
			t.Errorf("secondLargest(%v) failed, expected %d got %d", test.in, test.out, res)
		}
	}
}
