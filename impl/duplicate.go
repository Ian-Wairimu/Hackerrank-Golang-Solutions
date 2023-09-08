package main

import "fmt"

func main() {
	arr := []int{1, 3, 1, 2, 2, 4, 7, 5, 2, 6, 7, 3, 4, 8, 7}
	// n rep the length of the array
	n := len(arr)
	var frequency int
	for i := 0; i < n; i++ {
		frequency = 1
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				frequency++
				fmt.Println(frequency)
			}
		}
		if frequency > 1 {
			//fmt.Println(arr[i], frequency)
		}
	}
}
