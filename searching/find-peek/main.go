package main

import "fmt"

// Naive solution
// 0(N)
/* func peek(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	}

	if arr[0] > arr[1] {
		return 0
	}

	if arr[len(arr)-1] > arr[len(arr)-2] {
		return len(arr) - 1
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return arr[i]
		}
	}
} */

// 0(logN)
func peekIndex(arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if (mid == 0 || arr[mid-1] <= arr[mid]) &&
			(mid == len(arr)-1 || arr[mid+1] <= arr[mid]) {
			return mid
		}
		if mid > 0 && arr[mid-1] >= arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{10, 20, 15, 5, 23, 90, 67}
	fmt.Println(peekIndex(arr))
}
