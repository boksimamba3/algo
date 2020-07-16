package main

import "fmt"

func maxDiff(arr []int) int {
	maxDiff := 0

	for j := 1; j < len(arr); j++ {
		for i := 0; i < j; i++ {
			if arr[j]-arr[i] > maxDiff {
				maxDiff = arr[j] - arr[i]
			}
		}
	}

	return maxDiff
}

func main() {
	arr := []int{2, 3, 10, 6, 4, 17}
	fmt.Println(maxDiff(arr))
}
