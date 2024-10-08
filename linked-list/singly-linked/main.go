package main

import (
	"fmt"
)

type Element struct {
	next  *Element
	list  *List
	Value interface{}
}

func NewElement(v interface{}, list *List) *Element {
	return &Element{Value: v, list: list}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil {
		return p
	}
	return nil
}

type List struct {
	head *Element
	len  int
}

func New() *List {
	return (&List{}).Init()
}

func (l *List) Init() *List {
	l.head = nil
	l.len = 0
	return l
}

// IsEmpty checks if list is empty
func (l *List) IsEmpty() bool {
	return l.len == 0
}

// Len return number of elements in list
func (l *List) Len() int {
	return l.len
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.IsEmpty() {
		return nil
	}
	return l.head
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
	if l.IsEmpty() {
		return nil
	}

	p := l.head
	for p.next != nil {
		p = p.Next()
	}

	return p
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.head == nil {
		l.Init()
	}
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	e := NewElement(v, l)
	e.next = l.head
	l.head = e
	l.len++
	return e
}

// Push back inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	e := NewElement(v, l)
	if l.head == nil {
		e.next = l.head
		l.head = e
	} else {
		p := l.head
		for p.next != nil {
			p = p.Next()
		}
		e.next = p.next
		p.next = e
	}
	l.len++

	return e
}

// PopFront removes first element of the list and returns e
func (l *List) PopFront() *Element {
	if l.head == nil {
		return nil
	}
	e := l.head
	l.head = e.next
	e.next = nil
	l.len--

	return e
}

// PopBack removes last element of the list and returns e
func (l *List) PopBack() *Element {
	if l.head == nil {
		return nil
	}
	prev := l.head
	e := prev
	for e.next != nil {
		prev = e
		e = e.next
	}
	prev.next = nil
	l.len--
	return e
}

// InsertAt inserts value at postion n
func (l *List) InsertAt(n int, v interface{}) *Element {
	if n < 0 || n > l.len {
		return nil
	}
	e := NewElement(v, l)
	if n == 0 {
		e.next = l.head
		l.head = e
	} else {
		p := l.head
		for i := 1; i < n; i++ {
			p = p.next
		}
		e.next = p.next
		p.next = e
	}
	l.len++

	return e
}

// Reverse linked list
func (l *List) Reverse() {
	var previous *Element = nil
	var next *Element = nil
	current := l.head
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	l.head = previous
}

func (l *List) search(e *Element, v interface{}, i int) int {
	if e == nil {
		return -1
	}
	if e.Value == v {
		return i
	}
	return l.search(e.next, v, i+1)
}

// Search value in the list. If element is present in the list return his position otherwise -1.
// Recursive approach.
func (l *List) Search(v interface{}) int {
	i := 0
	e := l.head
	return l.search(e, v, i)
}

// Search value in the list. If element is present in the list return his position otherwise -1.
// Iterative approach.
/* func (l *List) Search(v interface{}) int {
	i := 0
	for p := l.head; p != nil; p = p.next {
		if p.Value == v {
			return i
		}
		i++
	}

	return -1
} */

func (l *List) Print() {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	fmt.Println("Before reverse")
	list.Print()
	// fmt.Println(list.Search(1))
	list.Reverse()
	fmt.Println("After reverse")
	list.Print()
}
