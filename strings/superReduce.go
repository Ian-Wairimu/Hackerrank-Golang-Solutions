package main

import (
	"fmt"
	"sort"
	"strings"
)

// "abcdefghijklmnopqrstuvwxyz"

func main() {
	str := "abba"
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] < chars[j]
	})

	str = string(chars)

	for i := 0; i < len(str); i += 2 {
		end := i + 2
		if end > len(str) {
			end = len(str)
		}
		if strings.Contains(str, str[i:end]) {
			trim := strings.Trim(str, str[i:end])
			fmt.Println(trim)
		}
	}
	//fmt.Println(str)
}

//
//func remove(ch string) {
//
//}
