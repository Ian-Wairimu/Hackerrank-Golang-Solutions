package main

import "fmt"

func maximizingXor(l int32, r int32) int32 {
	// Write your code here
	var xor int32
	var maxNum int32
	for i := l; i <= r; i++ {
		for j := l; j <= r; j++ {
			xor = i ^ j
			if xor > maxNum {
				maxNum = xor
			}
		}
	}
	return maxNum
}

func main() {
	var l, r int32
	l = 11
	r = 100
	n := maximizingXor(l, r)
	fmt.Println(n)
}
