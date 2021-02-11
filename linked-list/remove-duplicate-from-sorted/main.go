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

func removeDuplicatesFromList(list *List) {

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
	removeDuplicatesFromList(list)
	printList(list)

}
