package main

import "fmt"

func pairs(k int32, arr []int32) int32 {
	var count int32
	// Write your code here
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			sub := arr[i] - arr[j]
			if sub == k {
				count++
			}
		}
	}
	return count
}

func InsertionSort(n []int32) ([]int32, error) {
	// created an empty slice to store the already sorted data to the user
	var num []int32
	for i := 0; i < len(n); i++ {
		key := n[i]
		j := i - 1
		for j >= 0 && n[j] >= key {
			n[j+1] = n[j]
			j = j - 1
			n[j+1] = key
		}
	}
	// getting the already sorted data and appending it to already created slice
	num = append(num, n...)
	return num, nil
}

func main() {
	arr := []int32{1, 5, 3, 4, 2}
	var k int32
	k = 2
	sortedArr, _ := InsertionSort(arr)
	n := pairs(k, sortedArr)
	fmt.Println(n)
}
