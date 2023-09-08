package main

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var sum int64 = 0
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}
