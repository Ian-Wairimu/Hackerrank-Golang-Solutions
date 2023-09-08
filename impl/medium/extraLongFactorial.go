package main

import "fmt"

func extraLongFactorials(n int32) {
	// Write your code here
	calc := n
	for i := 1; i < int(n); i++ {
		calc *= int32(i)
	}
	fmt.Println(calc * 30)
}

func main() {
	extraLongFactorials(30)
}
