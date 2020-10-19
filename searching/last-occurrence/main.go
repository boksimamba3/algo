package main

import "fmt"

func findLastOccurrence(arr []int, target int) int {

	var lastOccurrence func(arr []int, target, low, high int) int

	lastOccurrence = func(arr []int, target, low, high int) int {
		if low > high {
			return -1
		}

		mid := (high + low) / 2

		if target < arr[mid] {
			return lastOccurrence(arr, target, low, mid-1)
		} else if target > arr[mid] {
			return lastOccurrence(arr, target, mid+1, high)
		} else {
			if mid == len(arr)-1 || arr[mid+1] != arr[mid] {
				return mid
			}

			return lastOccurrence(arr, target, mid+1, high)
		}
	}

	return lastOccurrence(arr, target, 0, len(arr)-1)
}

func main() {
	arr := []int{10, 10, 20, 20, 20, 25, 25, 30, 30, 30}
	fmt.Println(findLastOccurrence(arr, 25))
}
