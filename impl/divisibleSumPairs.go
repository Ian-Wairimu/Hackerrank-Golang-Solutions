package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count, i, j int32
	// Write your code here
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	n := 6
	k := 3
	arr := []int32{1, 3, 2, 6, 1, 2}
	num := divisibleSumPairs(int32(n), int32(k), arr)
	fmt.Println(num)
}
