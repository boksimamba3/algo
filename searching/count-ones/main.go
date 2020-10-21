package main

import "fmt"

// O(logN)
func countOnes(arr []int) int {

	var firstOccurrence func(arr []int, low, high int) int

	firstOccurrence = func(arr []int, low, high int) int {
		if low > high {
			return 0
		}

		mid := (low + high) / 2

		if arr[mid] < 1 {
			return firstOccurrence(arr, mid+1, high)
		}

		if mid == 0 || arr[mid-1] != arr[mid] {
			return len(arr) - mid
		}

		return firstOccurrence(arr, low, mid-1)
	}

	return firstOccurrence(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{0, 1, 1, 1}
	fmt.Println(countOnes(arr))
}
