package main

import "fmt"

// O(N)
func isEquilibriumPoint(arr []int) bool {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	leftSum := 0

	for i := 0; i < len(arr); i++ {
		if leftSum == sum-arr[i] {
			return true
		}

		leftSum += arr[i]
		sum -= arr[i]
	}

	return false
}

func main() {
	fmt.Println(isEquilibriumPoint([]int{3, 4, 8, -9, 20, 6}))
}
