package main

import "fmt"

func countOnes(arr []int) int {

	var firstOccurrence func(arr []int, target, low, high int) int

	firstOccurrence = func(arr []int, target, low, high int) int {
		if low > high {
			return -1
		}

		mid := (low + high) / 2

		if target < arr[mid] {
			return firstOccurrence(arr, target, low, mid-1)
		} else if target > arr[mid] {
			return firstOccurrence(arr, target, mid+1, high)
		} else {
			if mid == 0 || arr[mid-1] != arr[mid] {
				return mid
			}

			return firstOccurrence(arr, target, low, mid-1)
		}
	}

	indexOfFirst := firstOccurrence(arr, 1, 0, len(arr)-1)

	if indexOfFirst == -1 {
		return 0
	}

	return len(arr) - indexOfFirst
}

func main() {
	arr := []int{0, 0, 0, 1, 1, 1, 1}
	fmt.Println(countOnes(arr))
}
