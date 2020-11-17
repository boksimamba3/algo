package main

import "testing"

type TestData struct {
	in  int
	out bool
}

func TestPalindromeNumber(t *testing.T) {
	testData := []TestData{
		{10, false},
		{9559, true},
		{0, true},
		{78987, true},
		{8, true},
		{2456, false},
	}

	for _, test := range testData {
		res := isPalindrome(test.in)
		if test.out != res {
			t.Errorf("isPalindrome(%d) failed, expected %t got %t", test.in, test.out, res)
		}
	}
}
