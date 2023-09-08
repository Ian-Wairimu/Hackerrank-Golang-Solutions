package main

import (
	"fmt"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	// Write your code here
	var minArr []int32
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			minus := arr[i] - arr[j]
			if minus >= 0 {
				minArr = append(minArr, minus)
			}
		}
	}
	min := minArr[0]
	for _, v := range minArr {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	a := []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}
	n := minimumAbsoluteDifference(a)
	fmt.Println(n)
}
