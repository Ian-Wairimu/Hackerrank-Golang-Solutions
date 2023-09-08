package main

import (
	"fmt"
)

func minimumDistances(a []int32) int32 {
	// Write your code here
	var min, max int32
	var arr []int32
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				minus := int32(j - i)
				arr = append(arr, minus)
			}
		}
	}
	if arr != nil {
		min = arr[0]
	}
	for k := 0; k < len(arr); k++ {
		if arr[k] > max {
			max = arr[k]
		} else {
			if arr[k] < min {
				min = arr[k]
			}
		}
	}
	if min != 0 {
		return min
	}
	return -1
}

func main() {
	arr := []int32{1, 2, 3, 4, 10}
	//arr := []int32{1, 2, 3, 4, 10}
	n := minimumDistances(arr)
	fmt.Println(n)
}
