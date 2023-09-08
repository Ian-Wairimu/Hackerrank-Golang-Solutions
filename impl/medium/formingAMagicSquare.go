package main

import "fmt"

func main() {
	s := [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			fmt.Println(s[i][i])
			fmt.Println(s[j][j])
		}
	}
}
