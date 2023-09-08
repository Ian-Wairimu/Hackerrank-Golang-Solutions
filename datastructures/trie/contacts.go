package main

import (
	"fmt"
	"strings"
)

func contacts(queries [][]string) []int32 {
	// Write your code here
	var count, count2 int32
	n := len(queries)
	add := len("add")
	find := len("find")
	var arr []int32
	var st []string
	for i := 0; i < n; i++ {
		s := queries[i]
		for j := 0; j < len(s); j++ {
			str := s[j]
			if string(str[:add]) == "add" {
				strs := strings.TrimSpace(str[add:])
				st = append(st, strs)
			} else if string(str[:find]) == "find" {
				for k := 0; k < len(st); k++ {
					if strings.Contains(st[k], strings.TrimSpace(str[find:])) == true {
						count++
						continue
					}
					if strings.Contains(st[k], strings.TrimSpace(str[find:])) == false {
						count2 = 0
						break
					}

				}
			}
		}
	}
	arr = append(arr, count)
	arr = append(arr, count2)
	return arr
}

func main() {
	queries := [][]string{{"add hack"}, {"add hackerrank"}, {"find hac"}, {"find hak"}}
	//queries := [][]string{{"add ed"}, {"add, eddie"}, {"add, edward"}, {"find, ed"}, {"add, edwina"}, {"find, edw"}, {"find a"}}
	arr := contacts(queries)
	fmt.Println(arr)
}
