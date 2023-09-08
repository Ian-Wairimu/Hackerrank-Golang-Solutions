package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getMax(operations []string) []int32 {
	// Write your code here
	var arr []int32
	var newArr []int32
	for i := 0; i < len(operations); i++ {
		s := operations[i]
		if string(s[0]) == "1" {
			n, _ := strconv.Atoi(strings.TrimSpace(s[1:]))
			arr = append(arr, int32(n))
		} else if string(s[0]) == "2" {
			arr = arr[:len(arr)-1]
		} else if string(s[0]) == "3" {
			max := maximum(arr)
			newArr = append(newArr, max)
		}
	}
	return newArr
}
func maximum(arr []int32) int32 {
	var max int32
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func main() {
	str := []string{"1 97", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 91", "3"}
	arr := getMax(str)
	fmt.Println(arr)
}
