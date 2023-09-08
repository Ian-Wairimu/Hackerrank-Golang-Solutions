package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var hashmap = make(map[string]int32, len(h))
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	var arr []int32
	var max, total int32
	for i, v := range alphabet {
		hashmap[string(v)] = h[i]
	}
	for i := 0; i < len(word); i++ {
		arr = append(arr, hashmap[string(word[i])])
	}
	for j := 0; j < len(arr); j++ {
		if arr[j] > max {
			max = arr[j]
		}
	}
	total = max * int32(len(word))
	return total
}

func main() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	s := "zaba"
	n := designerPdfViewer(h, s)
	fmt.Println(n)
}
