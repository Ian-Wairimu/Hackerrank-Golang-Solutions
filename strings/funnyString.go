package main

import "fmt"

func funnyString(s string) string {
	// Write your code here
	var r []int32
	for i := 0; i < len(s); i++ {
		r = append(r, int32(s[i]))
	}
	for j := len(r) - 1; j >= 0; j-- {
		for k := len(r) - 1; k > 0; k-- {
			minus := r[j] - r[k]
			if minus >= 0 {
				fmt.Println(minus)
			}
		}
	}
	return ""
}

func main() {
	//str := "bcxz"
	str := "acxz"
	s := funnyString(str)
	fmt.Println(s)
}
