package main

import (
	"errors"
	"fmt"
)

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

func (bst *BST) Min() (int, error) {
	if bst.root == nil {
		return 0, errors.New("Tree is empty")
	}
	return min(bst.root).data, nil
}

func min(root *Node) *Node {
	if root.left == nil {
		return root
	}
	return min(root.left)
}

func (bst *BST) Max() (int, error) {
	if bst.root == nil {
		return 0, errors.New("Tree is empty")
	}
	return max(bst.root).data, nil
}

func max(root *Node) *Node {
	if root.right == nil {
		return root
	}
	return max(root.right)
}

func (bst *BST) Insert(data int) {
	bst.root = insert(bst.root, data)
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return NewNode(data)
	} else if data <= root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}
	return root
}

func (bst *BST) Delete(data int) {
	bst.root = delete(bst.root, data)
}

func delete(root *Node, data int) *Node {
	// Base case: Tree is empty
	if root == nil {
		return root
		// Recursively go left
	} else if data < root.data {
		root.left = delete(root.left, data)
		// Recursively go right
	} else if data > root.data {
		root.right = delete(root.right, data)
		// Node to be deleted
	} else {
		// Node with one or no child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		// Node with 2 children: Get inorder successor
		// smallest in the right subtree
		root.data = min(root.right).data
		// Delete the inorder successor
		root.right = delete(root.right, root.data)
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

func (bst *BST) InOrederPrint() {
	inorder(bst.root)
	fmt.Printf("\n")
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d ", root.data)
	inorder(root.right)
}

func main() {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(21)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(6)
	bst.Insert(11)
	fmt.Println("Search(30):", bst.Search(30))
	fmt.Println("Search(21):", bst.Search(21))
	fmt.Println("Delete(30)")
	//bst.Delete(30)
	fmt.Println("Delete(21)")
	bst.Delete(21)
	fmt.Println("Search(30):", bst.Search(30))
	fmt.Println("Search(21):", bst.Search(21))
	min, err := bst.Min()
	if err == nil {
		fmt.Println("Minimum of tree is:", min)
	}
	max, err := bst.Max()
	if err == nil {
		fmt.Println("Maximum of tree is:", max)
	}
	fmt.Println("*******PRINT*******")
	bst.InOrederPrint()
}
