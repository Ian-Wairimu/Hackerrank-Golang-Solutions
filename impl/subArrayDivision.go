package main

import (
	"fmt"
)

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	if len(s) == 1 {
		return 1
	}
	var count int32
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == d {
				arr := []int32{s[i], s[j]}
				if int32(len(arr)) == m {
					count++
				}
				break
			}
		}
	}
	if count != 0 {
		return count - 1
	}
	return 0
}

func main() {
	arr := []int32{1, 2, 1, 3, 2}
	m := 2
	d := 3
	n := birthday(arr, int32(d), int32(m))
	fmt.Println(n)
}
