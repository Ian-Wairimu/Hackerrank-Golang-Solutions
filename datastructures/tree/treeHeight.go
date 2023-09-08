package main

func Height(n *BinaryNode) int32 {
	if n == nil {
		return 0
	}
	x := Height(n.left)
	y := Height(n.right)
	lDepth := x
	rDepth := y
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}
