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

// Question: Sorted insert in the linked list.
// The linked list should remain sorted after insertion.

// Solution: O(N)
/*
	To insert elements in sorted order we need to find prev element
	that has value less then value we are inserting. After that we just
	hook up all the pointers like in ordinary linked list.
*/
func (l *List) SortedInsert(value int) {
	e := NewElement(value)
	if l.head == nil {
		e.next = l.head
		l.head = e
	} else {
		p := l.head
		for p.next != nil && p.next.value < value {
			p = p.next
		}
		e.next = p.next
		p.next = e
	}
}

func Print(l *List) {
	for p := l.head; p != nil; p = p.next {
		fmt.Println(p.value)
	}
}

func main() {
	list := NewList()
	list.SortedInsert(1)
	list.SortedInsert(3)
	list.SortedInsert(9)
	list.SortedInsert(4)
	list.SortedInsert(2)
	Print(list)
}
