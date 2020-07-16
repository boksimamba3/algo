package main

import "fmt"

// Element is called a leader of array if there
// is no element greater than it on the right side
// O(nˆ2)
/* func leader(arr []int) {
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
} */

// O(n)
func leader(arr []int) {
	n := len(arr)
	curr := arr[n-1]
	fmt.Printf("Leader [%d]\n", curr)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > curr {
			curr = arr[i]
			fmt.Printf("Leader [%d]\n", curr)
		}
	}
}

func main() {
	arr := []int{7, 10, 4, 3, 6, 5, 2}
	leader(arr)
}
