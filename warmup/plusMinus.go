package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var n, o, p float64
	var positive, negative, zero float64
	for i := 0; i < len(arr); i++ {
		length := len(arr)
		if arr[i] >= 1 {
			p++
			positive = p / float64(length)
		} else if arr[i] <= -1 {
			n++
			negative = n / float64(length)
		} else {
			o++
			zero = o / float64(length)
		}

	}
	fmt.Printf("%f\n%f\n%f", positive, negative, zero)
}
