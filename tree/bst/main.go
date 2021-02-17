package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Insert(data int) {
	bst.root = insert(bst.root, data)
}

func insert(root *Node, data int) *Node {
	if root == nil {
		root = NewNode(data)
	} else if data <= root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}
	return root
}

func (bst *BST) Search(data int) bool {
	return search(bst.root, data)
}

func search(root *Node, data int) bool {
	if root == nil {
		return false
	} else if data == root.data {
		return true
	} else if data <= root.data {
		return search(root.left, data)
	} else {
		return search(root.right, data)
	}
}

func main() {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(5)
	bst.Insert(15)
	fmt.Println("Search(15):", bst.Search(15))
	fmt.Println("Search(9):", bst.Search(9))
}
