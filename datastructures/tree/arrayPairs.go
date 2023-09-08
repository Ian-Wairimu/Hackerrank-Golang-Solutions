package main

import "fmt"

func solve(arr []int32) int64 {
	// Write your code here
	var newArr []int32
	var count int32
	for i := 1; i <= len(arr); i++ {
		if arr[i-1] != int32(i) {
			newArr = append(newArr, int32(i))
		}
	}
	newArr = append(newArr, arr...)
	for x := 0; x < len(newArr); x++ {
		for y := x + 1; y < len(newArr); y++ {
			c := newArr[x] * newArr[y]
			var max int32
			if c > max {
				max = c
			}
			if c <= max {
				count++
				break
			}
		}
	}
	return int64(count)
}

func removeDuplicates(s []int32) []int32 {
	bucket := make(map[int32]bool)
	var arr []int32
	for _, v := range s {
		if _, ok := bucket[v]; !ok {
			bucket[v] = true
			arr = append(arr, v)
		}
	}
	return arr
}

func main() {
	arr := []int32{1, 1, 2, 4, 2}
	n := solve(arr)
	fmt.Println(n)
}
