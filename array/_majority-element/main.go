package main

import "fmt"

func findMajority(arr []int) int {
	res := 0
	count := 1

	// Find a candidate
	for i := 1; i < len(arr); i++ {
		if arr[res] == arr[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			res = i
			count = 1
		}
	}

	// Check if the candidate is majority
	count = 0
	for i := 0; i < len(arr); i++ {
		if arr[res] == arr[i] {
			count++
		}
	}

	if count <= len(arr)/2 {
		res = -1
	}

	return res
}

func main() {
	fmt.Println(findMajority([]int{8, 3, 4, 8, 8}))
	fmt.Println(findMajority([]int{3, 7, 4, 7, 7, 5}))
	fmt.Println(findMajority([]int{3, 7, 4, 7, 7, 5, 7}))
}
