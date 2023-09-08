package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	// Write your code here
	var count int32
	for i := a; i <= b; i++ {
		sqrt := math.Sqrt(float64(i))
		if math.Mod(sqrt, 1) == 0 {
			count++
		}
	}
	return count
}

func main() {
	var a, b int32
	a = 3
	b = 9
	n := squares(a, b)
	fmt.Println(n)
}
