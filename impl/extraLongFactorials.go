package main

import "fmt"

func extraLongFactorials(n int32) int32 {
	if n == 0 {
		return 1
	}
	return n * extraLongFactorials(n-1)
}

func main() {
	n := extraLongFactorials(25)
	fmt.Println(n)
}
