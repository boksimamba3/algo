package main

import "fmt"

type Element struct {
	value int
	next  *Element
}

func NewElement(value int) *Element {
	return &Element{value: value}
}

type List struct {
	head *Element
}

func NewList() *List {
	return &List{}
}

func (l *List) Insert(value int) {
	e := NewElement(value)
	e.next = l.head
	l.head = e
}

// Question: Find the n-th node from the end of a given linked list.

// Solution: O(n) two iterations
/*
	First we need to find the length of the list. To do so we
	iterate trough the list and while doing so we increment the length.
	Once you found the length of the given linked list we need to
	iterate trough length - n, where n is position of of n-th element
	from back.
*/
/* func (l *List) FromBack(n int) int {
	len := 0
	p := l.head
	for p := l.head; p != nil; p = p.next {
		len++
	}
	if n > len {
		return -1
	}
	p = l.head
	for i := 0; i < len-n; i++ {
		p = p.next
	}
	return p.value
} */

// Solution: O(n) one iteration
/*
	Using two pointer approach we can solve this problem
	with one iteration. We move first pointer n positions so that
	difference between first and second pointer is n positions.
	Once we done that we move first and second pointer one step
	each iteration. When first pointer reaches end, second pointer
	will be at n-th position from back.
*/
func (l *List) FromBack(n int) int {
	if l.head == nil {
		return -1
	}
	first := l.head
	for i := 0; i < n; i++ {
		if first == nil {
			return -1
		}
		first = first.next
	}
	second := l.head
	for first != nil {
		first = first.next
		second = second.next
	}

	return second.value
}

func printList(l *List) {
	for p := l.head; p != nil; p = p.next {
		fmt.Println(p.value)
	}
}

func main() {
	list := NewList()
	list.Insert(5)
	list.Insert(4)
	list.Insert(3)
	list.Insert(2)
	list.Insert(1)
	printList(list)
	fmt.Printf("Nth from back [ %v ]\n", list.FromBack(1))  // 5
	fmt.Printf("Nth from back [ %v ]\n", list.FromBack(5))  // 1
	fmt.Printf("Nth from back [ %v ]\n", list.FromBack(10)) // -1
}
