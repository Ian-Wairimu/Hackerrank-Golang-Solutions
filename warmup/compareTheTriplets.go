package main

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var alice, bob int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	c := []int32{alice, bob}
	return c
}
