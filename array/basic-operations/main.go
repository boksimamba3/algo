package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// List is slice of ints
type List []int

// Search elements in list
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
func (list List) Reverse() {
	low := 0
	high := len(list) - 1

	for low < high {
		list[low], list[high] = list[high], list[low]
		low++
		high--
	}
}

// Circural access to list elements starting from position d
func (list List) Circural(d int) <-chan int {
	out := make(chan int)

	go func() {
		for i := d; i < len(list)+d; i++ {
			out <- list[i%len(list)]
		}
		close(out)
	}()

	return out
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
	fmt.Println("Position of number 4:", list.Search(4))
	fmt.Println("Is list sorted:", list.IsSorted())
	list.Reverse()
	fmt.Println("Reversed list:", list)
	list.Reverse()
	fmt.Println("Reversed back:", list)
	fmt.Println("Circural access to list:")
	for x := range list.Circural(2) {
		fmt.Println(x)
	}
}
