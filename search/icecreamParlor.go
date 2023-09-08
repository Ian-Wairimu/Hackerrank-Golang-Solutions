package main

import "fmt"

func icecreamParlor(m int32, arr []int32) []int32 {
	var x, y int32
	var newArr []int32
	// Write your code here
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == m {
				x = int32(i) + 1
				y = int32(j) + 1
			}
		}
	}
	newArr = append(newArr, x)
	newArr = append(newArr, y)
	return newArr
}

func main() {
	var m int32
	m = 4
	cost := []int32{1, 4, 5, 3, 2}
	//cost := []int32{2, 2, 4, 3}
	s := icecreamParlor(m, cost)
	fmt.Println(s)
}
