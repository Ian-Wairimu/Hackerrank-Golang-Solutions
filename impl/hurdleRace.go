package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var max, sub int32
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	sub = max - k
	if sub <= 0 {
		return 0
	}
	return sub
}

func main() {
	k := 4
	h := []int32{1, 6, 3, 5, 2}
	n := hurdleRace(int32(k), h)
	fmt.Println(n)
}
