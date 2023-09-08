package main

import "fmt"

func matchingStrings(stringList []string, queries []string) []int32 {
	// Write your code here
	visited := make(map[string]int32)
	var arr []int32
	for i := 0; i < len(stringList); i++ {
		for j := 0; j < len(queries); j++ {
			if stringList[i] == queries[j] {
				visited[stringList[i]]++
				break
			}
		}
	}
	for i := 0; i < len(queries); i++ {
		arr = append(arr, visited[queries[i]])
	}
	return arr
}

func main() {
	strList := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}
	arr := matchingStrings(strList, queries)
	fmt.Println(arr)
}
