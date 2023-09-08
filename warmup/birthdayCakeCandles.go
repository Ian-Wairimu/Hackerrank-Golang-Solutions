package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var sum int32
	var max = candles[0]
	for i := 0; i < len(candles); i++ {
		if max < candles[i] {
			max = candles[i]
		}
	}
	for j := 0; j < len(candles); j++ {
		if max == candles[j] {
			sum++
		}
	}
	return sum
}
