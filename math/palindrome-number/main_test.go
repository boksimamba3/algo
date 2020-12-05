package main

import "testing"

type TestCase struct {
	in  int
	out bool
}

func TestPalindromeNumber(t *testing.T) {
	cases := []TestCase{
		{10, false},
		{9559, true},
		{0, true},
		{78987, true},
		{8, true},
		{2456, false},
	}

	for _, test := range cases {
		res := isPalindrome(test.in)
		if test.out != res {
			t.Errorf("isPalindrome(%d) failed, expected %t got %t", test.in, test.out, res)
		}
	}
}
