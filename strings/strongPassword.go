package main

import (
	"fmt"
	"strings"
)

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var count int32
	numbers := "0123456789"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "!@#$%^&*()-+"
	if len(password) < 6 {
		count++
	}
	for i := 0; i < len(password); i++ {
		for j := 0; j < len(numbers); j++ {
			if strings.ContainsAny(string(password[i]), string(numbers[j])) == true {
				fmt.Println("yes")
				break
			}
		}
	}
	fmt.Println(numbers, lowercase, uppercase, specialCharacters)
	return count
}

func main() {
	psd := "Ab1"
	n := minimumNumber(int32(3), psd)
	fmt.Println(n)
}
