package main

import "fmt"

func utopianTree(n int32) int32 {
	// Write your code here
	var height []int32
	height = []int32{1, 2, 3, 6, 7, 14}
	tree := make(map[int32]int32, len(height))
	for i := 0; i <= 5; i++ {
		tree[int32(i)] = height[i]
	}
	for j := 0; j < len(tree); j++ {
		return tree[n]
	}
	return 0
}

func main() {
	n := utopianTree(4)
	fmt.Println(n)
}
