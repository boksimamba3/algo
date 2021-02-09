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

// Question: Reverse a linked list in a recursive manner.

// Solution: O(n)
func (l *List) Reverse() {
	var reverse func(head *Element) *Element
	reverse = func(head *Element) *Element {
		// Base case for empty list
		if head == nil {
			return head
		}
		// Base case when we reach end of the list
		if head.next == nil {
			return head
		}
		newHead := reverse(head.next)
		tail := head.next
		tail.next = head
		head.next = nil

		return newHead
	}

	l.head = reverse(l.head)
}

func Print(l *List) {
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
	Print(list)
	fmt.Println("Reversed")
	list.Reverse()
	Print(list)
}
