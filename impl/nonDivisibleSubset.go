package main

import "fmt"

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	var count int32
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			sum := s[i] + s[j]
			if sum%k != 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	//arr := []int32{1, 7, 2, 4}
	arr := []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}
	//k := int32(3)
	k := int32(7)
	n := nonDivisibleSubset(k, arr)
	fmt.Println(n)
}
