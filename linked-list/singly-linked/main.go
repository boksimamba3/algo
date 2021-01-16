package main

import "fmt"

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

// Len return number of elements in list
func (l *List) Len() int {
	return l.len
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
	if l.len == 0 {
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
func (l *List) PushBack(v interface{}) {
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
}

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
	list.Print()
}
