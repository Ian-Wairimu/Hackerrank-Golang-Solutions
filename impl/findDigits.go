package main

import (
	"fmt"
	"strconv"
)

func findDigits(n int32) int32 {
	// Write your code here
	strConv := strconv.Itoa(int(n))
	var arr []int32
	var count int32
	for i := 0; i < len(strConv); i += 1 {
		end := i + 1
		if end > len(strConv) {
			end = len(strConv)
		}
		inx, _ := strconv.Atoi(strConv[i:end])
		arr = append(arr, int32(inx))
	}
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 && n%arr[j] == 0 {
			count++
		}
	}
	return count
}

func main() {
	//num := 124
	num := 1012
	n := findDigits(int32(num))
	fmt.Println(n)
}
