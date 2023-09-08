package main

import "fmt"

func timeInWords(h int32, m int32) string {
	// Write your code here
	hour := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	//minute := []string{"one minute past", "ten minute past", "quarter past", "half past", "twenty minute to"}
	for i := 1; i <= len(hour); i++ {
		if int32(i) == h {
			fmt.Println(hour[i-1])
		}
	}
	return ""
}

func main() {
	timeInWords(5, 2)
}
