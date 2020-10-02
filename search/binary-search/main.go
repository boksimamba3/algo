package main

import "fmt"

func binarySearch(arr []int, x int) int {
	start := 0
	end := len(arr) - 1

	for end >= start {
		mid := (end + start) / 2

		if arr[mid] == x {
			return mid
		}

		if x < arr[mid] {
			end = mid - 1
		}

		if x > arr[mid] {
			start = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{10, 20, 30, 40, 60, 70}, 70))
}
