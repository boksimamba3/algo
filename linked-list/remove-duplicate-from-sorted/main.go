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
func removeDuplicatesFromList(list *List) {
	if list.head == nil {
		return
	}
	p := list.head
	for p.next != nil {
		next := p.next
		if next.value == p.value {
			p.next = next.next
			next.next = nil
		} else {
			p = p.next
		}
	}
}

func printList(l *List) {
	for p := l.head; p != nil; p = p.next {
		fmt.Println(p.value)
	}
}

func main() {
	list := NewList()
	list.Insert(5)
	list.Insert(5)
	list.Insert(4)
	list.Insert(3)
	list.Insert(3)
	list.Insert(2)
	list.Insert(1)
	list.Insert(1)
	list.Insert(1)
	printList(list)
	fmt.Println("Remove duplicates")
	removeDuplicatesFromList(list)
	printList(list)
}
