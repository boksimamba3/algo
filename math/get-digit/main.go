package main

import (
	"fmt"
	"math"
)

// Question: Find a number of digits in a number
// Example:  n: 9245, dp: 1 d: 5;  n: 9245, dp: 3 d: 2
// Note: In(n) > 0

// Solution: O(1)
/*
	Nth digit is reminder after dividing (n/10^(dp-1)) % 10
	n = 798, dp = 3
	d => (797 / 10 ^ 3 - 1) % 10 => (797 / 10^2) % 10 = 7

	n = 798, dp = 1
	d => (797 / 10 ^ 1 - 1) % 10 => (797 / 10^0) % 10 = 8
*/
func getDigit(n int, dp int) int {
	return (n / int(math.Pow(10, float64(dp-1)))) % 10
}

func main() {
	fmt.Println(getDigit(753, 1))
}
