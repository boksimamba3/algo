package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

// Naive solution O(n^2)
/* func getWater(arr []int) int {
	res := 0
	for i := 1; i < len(arr)-1; i++ {
		lMax := arr[i]
		for j := 0; j < i; j++ {
			lMax = max(lMax, arr[j])
		}

		rMax := arr[i]
		for j := i + 1; j < len(arr); j++ {
			rMax = max(rMax, arr[j])
		}

		res += min(lMax, rMax) - arr[i]
	}

	return res
} */

// Optimized solution
func getWater(arr []int) int {
	res := 0
	lMax := make([]int, len(arr))
	rMax := make([]int, len(arr))

	lMax[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		lMax[i] = max(arr[i], lMax[i-1])
	}

	rMax[len(arr)-1] = arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		rMax[i] = max(arr[i], rMax[i+1])
	}

	for i := 0; i < len(arr); i++ {
		res += min(lMax[i], rMax[i]) - arr[i]
	}

	return res
}

func main() {
	arr := []int{2, 1, 3, 1, 2, 0, 1}
	fmt.Println(getWater(arr))
}
