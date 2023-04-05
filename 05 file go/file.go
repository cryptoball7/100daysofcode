package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://gobyexample.com/writing-files

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {

	data := []byte("hello\ntext\n")
	err := os.WriteFile("test.txt", data, 0600)
	check(err)

	f, err := os.Create("data.txt")
	check(err)

	defer f.Close()
	
	data = []byte{115, 111, 109, 101, 10}
	bytesWritten, err := f.Write(data)
	check(err)
	fmt.Printf("Wrote %d bytes\n", bytesWritten)

	bytesWritten, err =f.WriteString("Hello again!\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten)
	
	f.Sync()

	w := bufio.NewWriter(f)
	bytesWritten, err = w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten)
	
	w.Flush()

}