package main

import "fmt"

// O(N)
func getSum(arr []int, start int, end int) int {
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}

	if start == 0 {
		return arr[end]
	}

	return arr[end] - arr[start-1]
}

func main() {
	fmt.Println(getSum([]int{2, 8, 3, 9, 6, 5, 4}, 0, 2))
	fmt.Println(getSum([]int{2, 8, 3, 9, 6, 5, 4}, 1, 3))
	fmt.Println(getSum([]int{2, 8, 3, 9, 6, 5, 4}, 2, 6))
}
