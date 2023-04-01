package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/fs"
)

func main() {
	fmt.Println("Hello, WOrld!")

	f,err := os.CreateTemp(filepath.Dir("tempfolder"), "tempfile.txt")

	fmt.Println(f)
	fmt.Println(err)

	fmt.Println(f.Name())

	info, err := os.Stat(f.Name())
	m := info.Mode()
	fmt.Println(err)
	fmt.Println(m)

	fmt.Println(fs.FileMode(int(0600)))

	os.Chmod(f.Name(), fs.FileMode(int(0600)))

	info2, err := os.Stat(f.Name())
	m = info2.Mode()
	fmt.Println(err)
	fmt.Println(m)
}