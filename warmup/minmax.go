package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var sum, min, max int32
	for _, v := range arr {
		sum += v
	}
	min = arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}
	totalX := sum - min
	totalY := sum - max
	fmt.Println(totalY, totalX)
}
