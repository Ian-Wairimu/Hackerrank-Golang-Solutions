package main

import "fmt"

func unexpectedDemand(arr []int32, k int32) int32 {
	var count int32
	for i := 0; i < len(arr); i++ {
		if arr[i] <= k {
			k -= arr[i]
			count++
		} else {
			break
		}
	}
	return count
}

func Insertion(n []int32) ([]int32, error) {
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
	// test case 1
	//m := 40
	m := 88492114
	//m := 3
	//m := 83178701
	//arr := []int{10, 30}
	//arr := []int{5, 4, 6}
	//arr := []int{21, 24}
	arr := []int32{94, 76532411, 4309294, 4173421, 3099762, 168026, 58739, 37856, 66235, 24078, 8222, 11597, 1092, 402, 83, 696, 56, 2, 8, 31, 25, 3, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	newArr, _ := Insertion(arr)
	num := unexpectedDemand(newArr, int32(m))
	fmt.Println(num)
}
