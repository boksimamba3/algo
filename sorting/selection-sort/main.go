package main

import "fmt"

// O(N^2)
func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := []int{4, 4, 1, 7, 2, 3, 0, 8}
	sort(arr)
	fmt.Println(arr)
}
