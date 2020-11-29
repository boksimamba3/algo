package main

import "fmt"

// Left rotate by one place
/* func leftRotate(arr []int) {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
} */

// Left rotate by n places Time O(N) Space O(d)
/* func leftRotate(arr []int, d int) {
	n := len(arr)

	tempArr := []int{}
	for i := 0; i < d; i++ {
		tempArr = append(tempArr, arr[i])
	}
	for i := d; i < n; i++ {
		arr[i-d] = arr[i]
	}
	for i := 0; i < len(tempArr); i++ {
		arr[n-d+i] = tempArr[i]
	}
} */

// Left rotate by n places Time O(N) Space O(1)
func leftRotate(arr []int, d int) {
	n := len(arr)
	reverse(arr, 0, d-1)
	reverse(arr, d, n-1)
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
	leftRotate(arr, 2)
	fmt.Println(arr)
}
