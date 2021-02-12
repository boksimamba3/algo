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
func merge(l1, l2 *List) *List {
	if l1.head == nil {
		return l2
	}
	if l2.head == nil {
		return l1
	}

	var newHead *Element
	p1 := l1.head
	p2 := l2.head
	p3 := newHead

	for p1 != nil && p2 != nil {
		if p1.value < p2.value {
			if newHead == nil {
				p3 = p1
				newHead = p3
			} else {
				p3.next = p1
				p3 = p1
			}
			p1 = p1.next
		} else {
			if newHead == nil {
				p3 = p2
				newHead = p3
			} else {
				p3.next = p2
				p3 = p2
			}
			p2 = p2.next
		}
	}

	for p1 != nil {
		p3.next = p1
		p3 = p3.next
		p1 = p1.next
	}

	for p2 != nil {
		p3.next = p2
		p3 = p3.next
		p2 = p2.next
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
	l3 := merge(l1, l2)
	printList(l3)
}
