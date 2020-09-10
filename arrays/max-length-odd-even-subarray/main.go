package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func remainder(x int) int {
	return x % 2
}

// Efficient solution 0(n)
func maxLenOddEvenSubarray(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	consecutiveMaxLength := 1
	consecutiveLength := 1
	previous := remainder(arr[0])

	for i := 1; i < len(arr); i++ {
		current := remainder(arr[i])
		if previous == current {
			consecutiveLength = 1
		} else {
			consecutiveLength++
		}
		previous = current
		consecutiveMaxLength = max(consecutiveMaxLength, consecutiveLength)
	}

	return consecutiveMaxLength
}

func main() {
	fmt.Println(maxLenOddEvenSubarray([]int{1, 2, 3, 4, 5, 6, 8, 10}))
	fmt.Println(maxLenOddEvenSubarray([]int{10, 12, 14, 7, 8}))
}
