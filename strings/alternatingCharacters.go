package main

import (
	"fmt"
)

func alternatingCharacters(s string) int32 {
	// Write your code here
	var count int32
	for i := 0; i < len(s)-1; i++ {
		fmt.Println(string(s[i]), string(s[i+1]))
		if s[i] == s[i+1] {
			count++
		}
	}
	return count
}

func main() {
	n := alternatingCharacters("AABAAB")
	fmt.Println(n)
}
