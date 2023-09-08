package main

import "fmt"

func sockMerchant(n int32, arr []int32) int32 {
	// Write your code here
	visited := make(map[int32]int32, len(arr))
	var count int32
	for i := 0; i < len(arr); i++ {
		visited[arr[i]]++
		if visited[arr[i]]%2 == 0 {
			count++
		}
		fmt.Println(visited[arr[i]])
	}
	return count
}
func main() {
	n := 9
	//arr := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	arr := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	num := sockMerchant(int32(n), arr)
	fmt.Println(num)
}
