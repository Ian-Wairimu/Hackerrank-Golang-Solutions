package main

import "fmt"

func maximumSumCrossing(a []int64, low, mid, high int64) int64 {
	var leftSum, rightSum, sum, calc int64
	for i := mid; i > low; i-- {
		sum = sum + a[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	sum = 0
	for j := mid + 1; j < high; j++ {
		sum = sum + a[j]
		if sum > rightSum {
			rightSum = sum
		}
	}
	calc = leftSum + rightSum
	return calc
}

func maximumSum(a []int64, m int64) int64 {
	// Write your code here
	var low, mid, max, high int64
	low = 0
	high = int64(len(a) - 1)
	mid = (low + high) / 2
	if low == high {
		return a[low]
	} else {
		n := maximumSumCrossing(a, low, mid, high)

		if n > max {
			max = n
		}
		return max % m
	}
}

func main() {
	arr := []int64{3, 3, 9, 9, 5}
	n := maximumSum(arr, int64(7))
	fmt.Println(n)
}
