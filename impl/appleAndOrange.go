package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	m := len(apples)
	n := len(oranges)
	var applesTotal, orangeTotal int32
	var applesCount, orangesCount int
	for i := 0; i < m; i++ {
		applesTotal = a + apples[i]
		if applesTotal >= s && applesTotal <= t {
			applesCount += 1
		}
	}
	for i := 0; i < n; i++ {
		orangeTotal = b + oranges[i]
		if orangeTotal >= s && orangeTotal <= t {
			orangesCount += 1
		}
	}
	fmt.Println(applesCount, orangesCount)
}

// s = 7
// t = 10
// appleTree = 4
// orange = 12
// there are m apples = 3
// there are n oranges =3
// apples = [2, 3, -4]
// oranges = [3, -2, -4]
// output 1 = apple, 2 oranges

func main() {
	var s, t, appleTree, orangeTree int32
	s = 7
	t = 10
	appleTree = 4
	orangeTree = 12
	apples := []int32{2, 3, -4}
	oranges := []int32{3, -2, -4}
	countApplesAndOranges(s, t, appleTree, orangeTree, apples, oranges)
}
