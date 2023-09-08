package main

import (
	"fmt"
)

func lonelyInteger(a []int32) int32 {
	// Write your code here
	visited := make(map[int32]int32, len(a))
	var num int32
	for i := 0; i < len(a); i++ {
		visited[a[i]]++
	}
	for j := 0; j < len(a); j++ {
		if visited[a[j]] == 1 {
			num = a[j]
		}
	}
	return num
}
func main() {
	n := []int32{34, 95, 34, 64, 45, 95, 16, 80, 80, 75, 3, 25, 75, 25, 31, 3, 64, 16, 31}
	l := lonelyInteger(n)
	fmt.Println(l)
}
