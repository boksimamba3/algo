package main

import "fmt"

// Question: Find a second largest element
// Example:
// arr: [10, 5, 8, 20]; index: 0,
// arr: [20, 10, 20, 8, 12]; index: 4
// arr: [10, 10, 10]; index: -1
// Note: arr[0], arr[1], arr[2]...arr[n] > 0

// Solution: O(N) 2 traversals
/*
	First find largest element in array. To find second
	larget we again traversal same array. This time we
	are looking for element different then first largest
	element.
*/
/* func secondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	// Find largest element
	largest := getLargest(arr)

	res := -1
	for i := 0; i < len(arr); i++ {
		// Current element is different from largest element
		if arr[i] != arr[largest] {
			// First element that is not equal to largest element
			if res == -1 {
				res = i
			// Current element is greater than previous
			// second largest element
			} else if arr[i] > arr[res] {
				res = i
			}
		}
	}

	return res
}

func getLargest(arr []int) int {
	if len(arr) < 1 {
		return -1
	}

	largest := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[largest] {
			largest = i
		}
	}

	return largest
} */

// Solution: O(N) 1 traversal
/*
	Keep track of largest and second largest element.
	First check if current element is larger than previous
	largest element, if so change second largest to
	previous largest, and curret largest is current element.
	If current element is not larger than current largest we need
	to check if it is different, if so check if we had
	previous second largest or current element is larger than previous
	second largest element. Otherwise ignore current element.
*/
func secondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	largest := 0
	res := -1

	for i := 1; i < len(arr); i++ {
		// Current element is larger than previous largest element
		if arr[i] > arr[largest] {
			res = largest
			largest = i
			// Current element if different than previous largest element
		} else if arr[i] != arr[largest] {
			// No previous largest element or current is greater then
			// previous largest
			if res == -1 || arr[i] > arr[res] {
				res = i
			}
		}
	}

	return res
}

func main() {
	fmt.Println(secondLargest([]int{10, 5, 8, 20}))
}
