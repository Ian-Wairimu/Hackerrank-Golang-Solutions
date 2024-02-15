package main

import "fmt"

func sumXor(n int64) int64 {
	// Write your code here
	var sum1, sum2, count, i, j int64
	var xorArr []int64
	var arr []int64

	for i = 0; i < n; i++ {
		sum1 = n ^ i
		xorArr = append(xorArr, int64(sum1))
	}
	for j = 0; j < n; j++ {
		sum2 = j + n
		arr = append(arr, int64(sum2))
	}
	x := len(xorArr)
	y := len(arr)
	if x == y {
		for i := 0; i < x; i++ {
			if xorArr[i] == arr[i] {
				count++
			}
		}
	}
	return count
}

func main() {
	n := 1111111113456
	sum := sumXor(int64(n))
	fmt.Println(sum)
}
