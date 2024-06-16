package main

import (
	"fmt"
)

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func newNode(key int) (n *Node) {
	n = &Node{
		key:    key,
		height: 1,
		left:   nil,
		right:  nil,
	}
	return
}

func insertNode(n *Node, key int) *Node {
	// recursively insert nodes
	switch {
	case n == nil:
		return newNode(key)
	case key < n.key:
		n.left = insertNode(n.left, key)
	case key > n.key:
		n.right = insertNode(n.right, key)
	// can't have the case of >= up above because it inserts itself
	default:
		return n
	}

	// if we haven't inserted a new node this code continues update the nodes height
	n.height = max(nodeHeight(n.left), nodeHeight(n.right)) + 1

	// check if the tree is unbalanced Positive is left heavy Negative right heavy
	weight := nodeHeight(n.left) - nodeHeight(n.right)
	switch {
	case weight > 1:
		if key < n.left.key {
			return rotateRight(n)
		} else if key > n.left.key {
			n.left = rotateLeft(n.left)
			return rotateRight(n)
		}
	case weight < -1:
		if key > n.right.key {
			return rotateLeft(n)
		} else if key < n.right.key {
			n.right = rotateRight(n.right)
			return rotateLeft(n)
		}
	}
	// its necessary to return the node here so that the recursive calls
	// don't delete nodes on inserts after second level
	return n
}

func rotateRight(n *Node) *Node {
	t1 := n.left
	t2 := t1.right
	t1.right = n
	n.left = t2
	t1.height = max(nodeHeight(n.left), nodeHeight(n.right)) + 1
	n.height = max(nodeHeight(n.left), nodeHeight(n.right)) + 1
	return t1
}

func rotateLeft(n *Node) *Node {
	t1 := n.right
	t2 := t1.left
	t1.left = n
	n.right = t2
	t1.height = max(nodeHeight(n.left), nodeHeight(n.right)) + 1
	n.height = max(nodeHeight(n.left), nodeHeight(n.right)) + 1
	return t1
}

func nodeHeight(n *Node) (h int) {
	if n != nil {
		h = n.height
	}
	return
}

func max(a, b int) (m int) {
	switch {
	case a > b:
		m = a
	default:
		m = b
	}
	return
}

func printNode(n *Node) {
	v := make([]*Node, 0)
	v = append(v, n)

	var bWT func(n *Node)
	bWT = func(n *Node) {
		if n == nil {
			return
		}
		if n.left != nil {
			v = append(v, n.left)
		}
		if n.right != nil {
			v = append(v, n.right)
		}
		bWT(n.left)
		bWT(n.right)
	}
	bWT(n)

	for _,k := range v {
	 fmt.Println(k.key)
	}
}

func main() {
	var t *Node
	s := []int{2, 1, 5, 4, 6, 3}
	for _, k := range s {
		t = insertNode(t, k)
	}
	printNode(t)
}
