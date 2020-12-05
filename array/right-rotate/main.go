package main

import "fmt"

// Question: Right rotate by d places
// Example:
// arr: [1, 2, 3, 4, 5]; rotatedOnePlace: [5, 1, 2, 3, 4]
// arr: [1, 2, 3, 4, 5]; rotatedTwoPlaces: [4, 5, 1, 2, 3]
// arr: [1, 2, 3, 4, 5]; rotatedThreePlaces: [3, 4, 5, 1, 2]

// Solution: 0(N) Space O(1)
/*
	First rotate array that starts at begining and
	ends at n-d-1, where d is number of places we want to
	right rotate array, and n is length of array.
	After that reverse the array that starts
	at n-d and ends at n-1.Once we rotated those two parts rotate the whole array.

	arr := {1, 2, 3, 4, 5} => {3, 4, 5, 1, 2}

	In = {1, 2, 3, 4, 5}, d = 2 Out: {3, 4, 5, 1, 2}
	I1: reverse(arr, 0, n-d) => {1, 2, 5, 4, 3}
	I2: reverse(arr, 0, n-d-1) => {2, 1, 5, 4, 3}
	I3: reverse(arr, 0, n-1) => {3, 4, 5, 1, 2}
*/
func rightRotate(arr []int, d int) {
	n := len(arr)
	reverse(arr, 0, n-d-1)
	reverse(arr, n-d, n-1)
	reverse(arr, 0, n-1)
}

func reverse(arr []int, low int, high int) {
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rightRotate(arr, 2)
	fmt.Println(arr)
}
