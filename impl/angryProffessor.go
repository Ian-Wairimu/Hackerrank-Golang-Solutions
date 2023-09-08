package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	count1 := 0
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			count1++
		}
	}
	if int32(count1) < k {
		return "YES"
	}
	return "NO"
}

func main() {
	a := []int32{-2, -1, 0, 1, 2}
	k := int32(3)
	str := angryProfessor(k, a)
	fmt.Println(str)
}
