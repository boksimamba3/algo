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

// 1 -> 2 -> 3 -> 4 -> 5
// 2 -> 4
func sortedMerge(l1, l2 *List) *List {
	if l1.head == nil {
		return l2
	}
	if l2.head == nil {
		return l1
	}

	var newHead *Element
	p1 := l1.head
	p2 := l2.head

	if p1.value < p2.value {
		newHead = p1
		p1 = p1.next
	} else {
		newHead = p2
		p1 = p2.next
	}
	tail := newHead

	for p1 != nil && p2 != nil {
		if p1.value < p2.value {
			tail.next = p1
			tail = p1
			p1 = p1.next
		} else {
			tail.next = p2
			tail = p2
			p2 = p2.next
		}
	}

	if p1 == nil {
		tail.next = p2
	} else {
		tail.next = p1
	}

	l1.head = newHead
	l2.head = nil

	return l1
}

func printList(l *List) {
	for p := l.head; p != nil; p = p.next {
		fmt.Println(p.value)
	}
}

func main() {
	l1 := NewList()
	l1.Insert(5)
	l1.Insert(4)
	l1.Insert(3)
	l1.Insert(2)
	l1.Insert(1)
	fmt.Println("First list")
	printList(l1)

	l2 := NewList()
	l2.Insert(4)
	l2.Insert(2)
	fmt.Println("Second list")
	printList(l2)

	fmt.Println("Merged list")
	l3 := sortedMerge(l1, l2)
	printList(l3)
}
