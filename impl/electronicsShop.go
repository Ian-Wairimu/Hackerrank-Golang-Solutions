package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	var arr []int32
	var max int32
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			sum := keyboards[i] + drives[j]
			if sum <= b {
				arr = append(arr, sum)
			}
		}
	}
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	if max != 0 {
		return max
	} else {
		return -1
	}
}

func main() {
	b := 60
	k := []int32{40, 50, 60}
	d := []int32{5, 8, 12}
	n := getMoneySpent(k, d, int32(b))
	fmt.Println(n)
}
