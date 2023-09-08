package main

import "fmt"

func InOrderTraversal(n *BinaryNode) {
	if n == nil {
		return
	}
	InOrderTraversal(n.left)
	fmt.Println(n.key)
	InOrderTraversal(n.right)
}
