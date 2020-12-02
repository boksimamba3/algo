package main

import "fmt"

// Question: Reverse an array
// Example:
// arr: [10, 5, 8, 20]; reversed: [20, 5, 8, 10]
// arr: [1, 2, 3, 4]; reversed: [4, 3, 2, 1]

// Solution: 0(N)
/*
	Swap left half with right half of the array.
*/
func reverse(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reverse(arr)
	fmt.Println(arr)
}
