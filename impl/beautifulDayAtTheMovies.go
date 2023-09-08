package main

import (
	"fmt"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var count int32
	for n := i; n <= j; n++ {
		minus := (n - reverse(n)) % k
		if minus == 0 {
			count++
		}
	}
	return count
}
func reverse(n int32) int32 {
	s := strconv.Itoa(int(n))
	rev := ""
	for i := len(s); i > 0; i-- {
		rev += string(s[i-1])
	}
	x, _ := strconv.Atoi(rev)
	return int32(x)
}
func main() {
	var i, j, k int32
	i = 13
	j = 45
	k = 3
	n := beautifulDays(i, j, k)
	fmt.Println(n)
}
