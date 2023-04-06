package main

import (
		"fmt"
		"os"
)

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