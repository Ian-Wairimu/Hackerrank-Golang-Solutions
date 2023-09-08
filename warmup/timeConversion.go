package main

import "strconv"

func timeConversion(s string) string {
	// Write your code here
	format := s[len(s)-2:]
	hours, _ := strconv.Atoi(s[:2])
	if format == "PM" && hours < 12 {
		hours = hours + 12
	} else if format == "AM" && hours == 12 {
		hours = hours - 12
	}
	newHours := strconv.Itoa(hours)
	converted := newHours + s[2:len(s)-2]
	if format == "AM" {
		converted = "0" + converted
	}
	return converted
}
