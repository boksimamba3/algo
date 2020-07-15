package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// List is slice of ints
type List []int

// Search finds elements in list
func (list List) Search(n int) int {
	for i := 0; i < len(list); i++ {
		if n == list[i] {
			return i
		}
	}

	return -1
}

// IsSorted checks if list is sorted
func (list List) IsSorted() bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}

	return true
}

// Reverse list
/* func (list List) Reverse() {
	n := len(list)
	for i := 0; i < n/2; i++ {
		list[i], list[n-1-i] = list[n-1-i], list[i]
	}
} */
// Reverse list
func (list List) Reverse() {
	low := 0
	high := len(list) - 1

	for low < high {
		list[low], list[high] = list[high], list[low]
		low++
		high--
	}
}

// String formats output
func (list List) String() string {
	var b bytes.Buffer
	l := len(list)
	for i, v := range list {
		b.WriteString(strconv.Itoa(v))
		if i != l-1 {
			b.WriteString(",")
		}
	}

	return b.String()
}

func main() {
	list := List{1, 2, 3, 4, 5}
	//fmt.Println(list.Search(4))
	//fmt.Println(list.IsSorted())
	list.Reverse()
	fmt.Println(list)
	list.Reverse()
	fmt.Println(list)
}
