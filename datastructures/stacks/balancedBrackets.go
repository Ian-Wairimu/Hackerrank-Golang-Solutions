package main

import "fmt"

func isBalanced(s string) string {
	// Write your code here
	//var revC string
	n := len(s)
	// f to represent the first
	f := s[:n/2]
	// c to represent the second
	c := s[n/2:]

	for i := 0; i < len(f); i++ {
		fmt.Println(f[i])
	}
	for j := 0; j < len(c); j++ {
		fmt.Println(c[j])
	}
	return ""
}

func main() {
	//str := "{[()]}"
	str := "{[(])}"
	isBalanced(str)
}
