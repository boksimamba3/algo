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

// Question: Find middle element of linked list.

// Solution: O(N)
/*
	We have two pointer. One slow and one fast. Slow moves
	one element(node) at time while fast moves two elements(nodes).
	When fast reaches end slow will reach the middle.
*/
func (l *List) GetMiddle() int {
	if l.head == nil {
		return -1
	}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.value
}

func printList(l *List) {
	for p := l.head; p != nil; p = p.next {
		fmt.Println(p.value)
	}
}

func main() {
	list := NewList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(9)
	list.Insert(4)
	list.Insert(2)
	list.Insert(5)
	printList(list)
	fmt.Printf("Middle element [ %v ]", list.GetMiddle())
}
