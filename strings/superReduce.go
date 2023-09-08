package main

import (
	"fmt"
	"strings"
)

// "abcdefghijklmnopqrstuvwxyz"

func main() {
	s := "aaabccddd"
	remove(s)
	newString := ""
	visited := make(map[string]int32)
	for _, v := range s {
		visited[string(v)]++
	}
	for i := 0; i < len(s); i++ {
		if visited[string(s[i])] > 1 {
			ch := strings.Trim(s, string(s[i]))
			continue
			newString += ch
		}
	}
	fmt.Println(newString)
}
func remove(ch string) {
	str := "a"
	for i := 0; i < len(ch); i++ {
		if string(ch[i]) == str {

		}
	}
}
