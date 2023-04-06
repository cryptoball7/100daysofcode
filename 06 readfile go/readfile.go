package main

import (
		"fmt"
		"os"
)

// https://gobyexample.com/reading-files

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("data.txt")
	check(err)
	fmt.Print(string(data))
}