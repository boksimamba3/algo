package main

import "fmt"

// Remove duplicates from sorted array (NAIVE)
// T O(N)
// S O(N)
/* func removeDuplicates(arr []int) []int {
	unique := []int{}
	if len(arr) == 0 {
		return unique
	}
	curr := arr[0]
	unique = append(unique, arr[0])
	for i := 1; i < len(arr); i++ {
		if curr != arr[i] {
			unique = append(unique, arr[i])
			curr = arr[i]
		}
	}
	return unique
} */

// Remove duplicates from sorted array (no extra space)
func removeDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	idx := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[idx-1] {
			arr[idx] = arr[i]
			idx++
		}
	}
	fmt.Println(arr)
	return arr[0:idx]
}

func main() {
	arr := []int{10, 20, 20, 30, 30, 30, 30, 40, 40, 50, 60}
	unique := removeDuplicates(arr)
	fmt.Println(unique)
}
