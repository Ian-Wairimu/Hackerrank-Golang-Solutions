package main

import (
	"fmt"
)

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var num int32
	visited := make(map[int32]int32, len(arr))
	var n int32
	n = 1
	for i := 1; i < len(arr); i++ {
		visited[arr[i]]++
		if visited[arr[i]] > n {
			num = arr[i-1]
		}
	}
	return num
}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	//arr := []int32{1, 4, 4, 4, 5, 3}
	n := migratoryBirds(arr)
	fmt.Println(n)
}
