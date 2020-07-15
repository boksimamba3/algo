package main

import "fmt"

// Left rotate by one place
/* func leftRotate(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp

	return arr
} */

// Left rotate by n places
func leftRotate(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

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

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(leftRotate(arr, 4))
}
