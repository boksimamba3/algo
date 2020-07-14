package main

import "fmt"

// Remove duplicates from sorted array (NAIVE)
func removeDuplicates(arr []int) []int {
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
}

func main() {
	arr := []int{10, 20, 20, 30, 30, 30, 30, 40, 40, 50}
	unique := removeDuplicates(arr)
	fmt.Println(unique)
}
