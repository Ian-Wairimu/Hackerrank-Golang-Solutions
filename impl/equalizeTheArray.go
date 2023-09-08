package main

import "fmt"

func equalizeArray(arr []int32) int32 {
	// Write your code here
	var max int32
	visited := make(map[int32]int32, len(arr))
	for _, v := range arr {
		visited[v]++
	}
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] > max {
			max = visited[arr[i]]
		}
	}
	return int32(len(arr)) - max
}

func main() {
	arr := []int32{3, 3, 2, 1, 3}
	n := equalizeArray(arr)
	fmt.Println(n)
}
