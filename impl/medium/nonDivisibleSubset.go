package main

import "fmt"

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	arr := []int32{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			sum := s[i] + s[j]
			if sum%k != 0 {
				arr = append(arr, s[i])
				arr = append(arr, s[j])
				break
			}
		}
	}
	return int32(len(arr) - 1)
}

func main() {
	arr := []int32{1, 7, 2, 4}
	a := 3
	n := nonDivisibleSubset(int32(a), arr)
	fmt.Println(n)
}
