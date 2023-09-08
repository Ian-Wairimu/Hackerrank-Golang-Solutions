package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var x, y int32
	var o, p int32
	x = a[len(a)-1]
	y = b[0]
	var arr []int32
	for i := x; i <= y; i++ {
		if i%x == 0 {
			arr = append(arr, i)
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(a); j++ {
			if arr[i]%a[j] != 0 {
				o++
				break
			}
		}
		for k := 0; k < len(b); k++ {
			if b[k]%arr[i] != 0 {
				p++
				break
			}
		}
	}
	if o > p {
		return int32(len(arr)) - o
	}
	return int32(len(arr)) - p
}

func main() {
	//a := []int32{2, 4}
	//b := []int32{16, 32, 96}
	a := []int32{3, 4}
	b := []int32{24, 48}
	n := getTotalX(a, b)
	fmt.Println(n)
}
