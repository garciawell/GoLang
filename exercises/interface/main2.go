package main

import (
	"fmt"
	"io"
	"os"
)

type size struct{}

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
