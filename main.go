package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	palind := x
	convAscii := strconv.Itoa(x)
	newChar := ""
	for i := len(convAscii); i > 0; i-- {
		newChar += string(convAscii[i-1])
	}
	convInt, _ := strconv.Atoi(newChar)
	fmt.Println(convInt)
	if convInt == palind {
		return true
	}
	return false
}
func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
		switch string(s[i]) {
		case "I":
			sum = sum + 1
			continue
		case "V":
			sum = sum + 5
			continue
		case "X":
			sum = sum + 10
			continue
		case "L":
			sum = sum + 50
			continue
		case "C":
			sum = sum + 100
			continue
		case "D":
			sum = sum + 500
			continue
		case "M":
			sum = sum + 1000
			continue
		default:
			sum = 0
		}
	}
	return sum
}

func main() {
	//n := isPalindrome(121)
	//fmt.Println(n)

	//x := romanToInt("MCMXCIV")
	//fmt.Println(x)
	//fmt.Println(bits.Add(10, 10, 0))
}
