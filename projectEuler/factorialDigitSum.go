package main

import "fmt"

func main() {
	n := 6
	mult := 0
	for i := 0; i < n; i++ {
		mult = n * (n - 1)
	}
	fmt.Println(mult)
}
