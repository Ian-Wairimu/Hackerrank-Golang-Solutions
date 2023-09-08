package main

import (
	"fmt"
	"strconv"
)

func marsExploration(s string) int32 {
	n := 3
	var result []string
	var count int32
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	fmt.Println(result)
	for i := 0; i < len(result); i++ {
		if result[i] != "SOS" {
			count++
		} else {
			count = 0
		}
	}
	return count
}

//func main() {
//	//s := "SOSSOT"
//	//s := "QQQ"
//	s := "OOSDSSOSOSWEWSOSOSOSOSOSOSSSSOSOSOSS"
//	fmt.Println(len(s))
//	//s := "SOSSPSSQSSOR"
//	//s := "SOSSOSSOS"
//	n := marsExploration(s)
//	fmt.Println(n)
//}

func main() {
	num := 1234
	str := strconv.Itoa(num)
	n := 1
	arr := []string{}
	for i := 0; i < len(str); i += n {
		end := i + n
		if end > len(str) {
			end = len(str)
		}
		arr = append(arr, str[i:end])
	}
	fmt.Println(arr)
}
