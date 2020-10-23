package main

import "fmt"

func binarySearch(arr []int, x, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == x {
		return mid
	} else if arr[mid] > x {
		return binarySearch(arr, x, low, mid-1)
	} else {
		return binarySearch(arr, x, mid+1, high)
	}
}

// NOTE THAT THIS FUNCTION ASSUMES arr[] TO BE OF INFINITE SIZE
// THEREFORE, THERE IS NO INDEX OUT OF BOUND CHECKING
// O(logN)
func search(arr []int, x int) int {
	if arr[0] == x {
		return 0
	}

	i := 1

	for arr[i] < x {
		i = i * 2
	}

	if arr[i] == x {
		return i
	}

	return binarySearch(arr, x, i/2+1, i-1)
}

func main() {
	arr := []int{1, 20, 22, 23, 50, 100, 1000, 2000, 4000}
	fmt.Println(search(arr, 100))
}
