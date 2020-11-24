package main

import "fmt"

// Question: Find a number of trailing zeros in factorial
// Example:  In: 5, Out: 1; In:10 , Out: 2
// Note: Number can overflow

// Solution: Naive O(N)
// Problem: Number overflow
/*
	Compute factorial number and while modulo of 10 equal
	to zero increment the count of trailing zeros.
*/
/* func countZeros(n int) int {
	fact := 1
	// Compute factorial number
	for i := 2; i <= n; i++ {
		fact = fact * i
	}
	// Count trailing zeros
	count := 0
	for fact%10 == 0 {
		count++
		fact = fact / 10
	}

	return count

} */

// Solution: Optimized O(logN) logBase=5
// Note: Prime factorization: which prime numbers multiply
// together to make the original number.
/*
	Count how many twos and fives we have in prime factorization
	of give factorial. Pair of two and five forms one trailing zero.
	We can only count number of five since number of fives will be
	allways less then number of twos.
	1x2x3x4x5x6x7x8x9x10...n
	Trailing zero count => floor(n/5)+floor(n/25)+floor(n/125)...
	In = 251, Out = 62
	n = 251, res = 0
	I1: i = 5, res => 0 + 251/5 => 50
	I2: i = 25, res => 50 + 251/25 => 60
	I3: i = 125, res => 60 + 251/125 => 62
	I4: i = 625, exit
*/
func countZeros(n int) int {
	res := 0
	for i := 5; i <= n; i = i * 5 {
		res = res + n/i
	}

	return res
}

func main() {
	fmt.Println(countZeros(251))
}
