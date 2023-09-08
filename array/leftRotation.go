package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	x := arr[:d]
	y := arr[d:]
	var a []int32
	a = append(a, y...)
	a = append(a, x...)
	return a
}

func main() {
	num := []int32{1, 2, 3, 4, 5}
	arr := rotateLeft(int32(4), num)
	fmt.Println(arr)
}
