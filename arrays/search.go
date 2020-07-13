package main

import "fmt"

// List of 5 elements
type List [5]int

// Search finds elements in list
func (list List) Search(n int) int {
	for i := 0; i < len(list); i++ {
		if n == list[i] {
			return i
		}
	}

	return -1
}

/* func search(arr [5]int, n int) int {
	for i := 0; i < len(arr); i++ {
		if n == arr[i] {
			return i
		}
	}

	return -1
}
*/
func main() {
	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	/* arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(search(arr, 3)) */
	list := List{1, 2, 3, 4, 5}
	fmt.Println(list.Search(3))
}
