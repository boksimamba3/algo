package main

import "fmt"

func leader(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		isLeader := true
		for j := i + 1; j < n; j++ {
			if arr[i] < arr[j] {
				isLeader = false
				break
			}
		}
		if isLeader {
			fmt.Printf("Leader [%d]\n", arr[i])
		}
	}
}

func main() {
	arr := []int{7, 10, 4, 3, 6, 5, 2}
	leader(arr)
}
