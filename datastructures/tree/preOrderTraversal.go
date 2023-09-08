package main

import (
	"fmt"
)

type BinaryNode struct {
	key   int32
	left  *BinaryNode
	right *BinaryNode
}
type BinaryTree struct {
	root *BinaryNode
}

func PreOrder(n *BinaryNode) {
	if n == nil {
		return
	}
	fmt.Print(fmt.Sprintf("%d ", n.key))
	PreOrder(n.left)
	PreOrder(n.right)
}

func (t *BinaryTree) Insert(n *BinaryNode) {
	var y *BinaryNode
	x := t.root
	for x != nil {
		y = x
		if n.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	if y == nil {
		t.root = n
	} else if n.key < y.key {
		y.left = n
	} else {
		y.right = n
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var data int
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&data)
		arr = append(arr, data)
	}
	tree := &BinaryTree{}
	for i := 0; i < len(arr); i++ {
		node := &BinaryNode{}
		node.key = int32(arr[i])
		tree.Insert(node)
	}
	PreOrder(tree.root)
}
