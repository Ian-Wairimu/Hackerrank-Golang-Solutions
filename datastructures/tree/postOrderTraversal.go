package main

import "fmt"

func PostOrderTraversal(n *BinaryNode) {
	if n == nil {
		return
	}
	PostOrderTraversal(n.left)
	PostOrderTraversal(n.right)
	fmt.Println(n.key)
}
