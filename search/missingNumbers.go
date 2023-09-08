package main

import (
	"fmt"
)

func missingNumbers(arr []int32, brr []int32) []int32 {
	// Write your code here
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(brr); j++ {
			if arr[i] != brr[j] {
				fmt.Println(brr[j])
				break
			}
		}
	}
	return nil
}
func main() {
	arr := []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
	brr := []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}
	asc := missingNumbers(arr, brr)
	//204 205 206
	fmt.Println(asc)
}
