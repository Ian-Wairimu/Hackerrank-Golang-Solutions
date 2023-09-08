package main

import (
	"fmt"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	// Write your code here
	if len(s) <= 1 {
		return n
	}
	var count int64
	iteratedString := strings.Repeat(s, int(n))
	if int64(len(iteratedString)) > n {
		iteratedString = iteratedString[:n]
	}
	for j := 0; j < len(iteratedString); j++ {
		if string(iteratedString[j]) == "a" {
			count++
		}
	}
	return count
}

func main() {
	//s := "a"
	s := "aba"
	//s := "epsxyyflvrrrxzvnoenvpegvuonodjoxfwdmcvwctmekpsnamchznsoxaklzjgrqruyzavshfbmuhdwwmpbkwcuomqhiyvuztwvq"
	//n := 1000000000000
	n := 10
	//n := 549382313570
	c := repeatedString(s, int64(n))
	fmt.Println(c)
}
