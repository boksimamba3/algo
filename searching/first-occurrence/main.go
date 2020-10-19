package main

import "fmt"

func findFirstOccurrence(arr []int, target int) int {

	var firstOccurrence func(arr []int, target, low, high int) int

	firstOccurrence = func(arr []int, target, low, high int) int {
		if low > high {
			return -1
		}

		mid := (high + low) / 2

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

	return firstOccurrence(arr, target, 0, len(arr)-1)
}

func main() {
	arr := []int{10, 10, 20, 20, 20, 25, 25, 30, 30, 30}
	fmt.Println(findFirstOccurrence(arr, 25))
}
