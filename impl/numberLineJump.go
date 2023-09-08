package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	for x1 < x2 {
		x1 += v1
		x2 += v2
		if x1 == x2 {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	//s := kangaroo(0, 3, 4, 2)
	//s := kangaroo(0, 2, 5, 3)
	s := kangaroo(4523, 8092, 9419, 8076)
	fmt.Println(s)
}
