package main

import "fmt"

func sansaXor(arr []int32) int32 {
	var i, j, total int32
	n := len(arr)
	var sumArr []int32
	var newArr []int32
	var sumAll int32

	for i = 1; i < int32(n); i++ {
		var sum = (i) ^ arr[i]
		sumArr = append(sumArr, sum)
	}
	for j = 0; j < int32(n); j++ {
		sumAll ^= arr[j]
		fmt.Println(sumAll)
		sumArr = append(sumArr, sumAll)
	}
	newArr = append(newArr, arr...)
	newArr = append(newArr, sumArr...)
	fmt.Println(newArr)
	for _, v := range newArr {
		total ^= v
	}
	return total
}

func main() {
	arr := []int32{1, 2, 3}
	total := sansaXor(arr)
	fmt.Println(total)

}
