package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 200)
	count, _ := file.Read(data)
	fmt.Println(count, string(data[:count]))
}
