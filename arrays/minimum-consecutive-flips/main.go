package main

import "fmt"

func minConsecutiveFlips(arr []int) {

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if arr[i] != arr[0] {
				fmt.Printf("From %d to", i)
			} else {
				fmt.Printf(" %d\n", i-1)
			}
		}
	}
	if arr[len(arr)-1] != arr[0] {
		fmt.Printf(" %d\n", len(arr)-1)
	}
}

func minConsecutiveFlipsCount(arr []int) int {
	flips := 0
	consecutive := false
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[0] {
			if !consecutive {
				flips++
			}
			consecutive = true
		} else {
			consecutive = false
		}
	}

	return flips
}

func main() {
	minConsecutiveFlips([]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0})
	fmt.Println("Min consecutive flips count: ", minConsecutiveFlipsCount([]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0}))
}
