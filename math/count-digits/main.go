package main

import (
	"fmt"
	"math"
)

// Question: Find a number of digits in a number
// Example:  In: 9245, Out: 4;  In: 43256, Out: 5; In: 7, Out: 1;
// Note: In(x) > 0

// Solution: Naive O(d) d = digits
/*
	Divide number by 10 until is greater then 0.
	In = 798, Out = 3
	x = 798, digits = 0
	I1: x => 798/10 = 79 digits = 1
	I2: x => 79/10 = 7 digits = 2
	I3: x => 7/10 = 0 digits = 3
*/
/* func countDigits(x int) int {
	if x <= 0 {
		return 0
	}

	digits := 0
	for x > 0 {
		digits++
		x = x / 10
	}

	return digits
} */

// Solution: Optimized O(1)
/*
	Using upper bound of log10. We know that
	log10(10^1)=1, log10(10^2)=2, log10(10^3)=3...
	and our number is between 10^2 and 10^3 or 100 and 1000.
	log = log10(798) = 2.90200289135, digits = floor(log)+1
	In: 798, Out: 3,
	x = 798, digits = log10(in) + 1
*/
func countDigits(x int) int {
	if x <= 0 {
		return 0
	}

	log := math.Log10(float64(x))
	return int(math.Floor(log) + 1)
}

func main() {
	fmt.Println(countDigits(9567))
}
