package main

import "fmt"

// Question: Find if a number is palindrome. Number is palindrome if
// it reads the same backward as forward.
// Example:  In: 78987, Out: true;  In: 8, Out: true; In: 254, Out: false;
// Note: In(n) >= 0

// Solution: Naive N(d) d = digits
/*
	Last digit n % 10. Remaining n/10.
	Reverse number is rev = rev * 10 + lastDigit.
	Compare if reverse number is equal to original number.
	In = 787, Out = true
	n = 787, rev = 0, temp = 787
	I1: lastDigit = 7, temp => 787/10 = 78, rev => 0 * 10 + 7 = 7
	I2: lastDigit = 8, temp => 78/10 = 7, rev => 7 * 10 + 8 = 78
	I3: lastDigit = 7, temp => 7/10 = 0, rev => 78 * 10 + 8 = 787
*/
func isPalindrome(n int) bool {
	rev := 0
	temp := n
	for temp != 0 {
		lastDigit := temp % 10
		rev = rev*10 + lastDigit
		temp = temp / 10
	}

	return rev == n
}

func main() {
	fmt.Println(isPalindrome(78987))
}
