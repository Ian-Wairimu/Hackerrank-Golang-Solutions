package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32
	for i := 0; i < len(bill); i++ {
		if int32(i) == k {
			continue
		}
		sum += bill[i]
	}
	sub := b - (sum / 2)
	if sub == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(sub)
	}
}

func main() {
	arr := []int32{3, 10, 2, 9}
	k := 1
	b := 7
	bonAppetit(arr, int32(k), int32(b))
}
