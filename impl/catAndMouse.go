package main

import "fmt"

func catAndMouse(x int32, y int32, z int32) string {
	word := ""
	subA := z - x
	subB := z - y
	if subA < 0 {
		subA = subA * -1
	}
	if subB < 0 {
		subB = subB * -1
	}
	if subA == subB {
		word += "MOUSE C"
	} else if subA > subB {
		word += "CAT B"
	} else if subB > subA {
		word += "CAT A"
	}
	return word
}

func main() {
	var x, y, z int32
	x = 1
	y = 2
	z = 3
	s := catAndMouse(x, y, z)
	fmt.Println(s)
}
