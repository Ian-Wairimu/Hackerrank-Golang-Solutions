package main

import "fmt"

func theGreatXor(x int64) int64 {
	// Write your code here
	var i int64
	count := 0
	for i = 1; i <= x; i++ {
		sum := i ^ x
		if sum > x {
			count++
		}
	}
	return int64(count)
}

func main() {
	x := 1303206072
	n := theGreatXor(int64(x))
	fmt.Println(n)
}
