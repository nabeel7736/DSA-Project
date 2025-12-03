package bst

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(value int) {
	bst.root = insertHelper(bst.root, value)
}

func insertHelper(n *Node, value int) *Node {
	if n == nil {
		return &Node{value: value}
	}
	if value < n.value {
		n.left = insertHelper(n.left, value)
	} else if value > n.value {
		n.right = insertHelper(n.right, value)
	}
	return n
}

func (bst *BST) Search(value int) bool {
	return searchHelper(bst.root, value)
}

func searchHelper(n *Node, value int) bool {
	if n == nil {
		return false
	}
	if value == n.value {
		return true
	}
	if value < n.value {
		return searchHelper(n.left, value)
	}
	return searchHelper(n.right, value)
}

func findMin(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (bst *BST) Delete(value int) {
	bst.root = deleteHelper(bst.root, value)
}

func deleteHelper(n *Node, value int) *Node {
	if n == nil {
		return nil
	}
	if value < n.value {
		n.left = deleteHelper(n.left, value)
	} else if value > n.value {
		n.right = deleteHelper(n.right, value)
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		minNode := findMin(n.right)
		n.value = minNode.value
		n.right = deleteHelper(n.right, minNode.value)
	}
	return n
}

func (bst *BST) Inorder() {
	inorderHelper(bst.root)
	fmt.Println()
}

func inorderHelper(n *Node) {
	if n != nil {
		inorderHelper(n.left)
		fmt.Print(n.value, " ")
		inorderHelper(n.right)
	}
}
