package main

import "fmt"

// O(N^2)
func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{4, 4, 1, 8, 2, 3, 0, 4}
	sort(arr)
	fmt.Println(arr)
}
