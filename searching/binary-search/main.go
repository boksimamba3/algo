package main

import "fmt"

/* func binarySearchR(arr []int, target int) int {

	var binarySearch func(arr []int, target, low, high int) int

	binarySearch = func(arr []int, target, low, high int) int {
		if low > high {
			return -1
		}

		mid := (high + low) / 2

		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			return binarySearch(arr, target, low, mid-1)
		}

		if target > arr[mid] {
			return binarySearch(arr, target, mid+1, high)
		}

		return -1
	}

	return binarySearch(arr, target, 0, len(arr)-1)

}
*/

// O(logN)
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for high >= low {
		mid := (high + low) / 2

		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			high = mid - 1
		}

		if target > arr[mid] {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{10, 20, 30, 40, 60, 70}, 70))
}
