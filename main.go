package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 0. check any args preventing program crash
	if len(os.Args) <= 1 {
		fmt.Println("Please pass an argument")
		os.Exit(1)
	}

	// 1. get the file name
	// os.Args  https://golang.org/pkg/os/#pkg-variables
	fileName := os.Args[1]
	fmt.Println("I got this file name:", fileName)

	// 2. read and print it by using io.Copy
	// os.Open  https://golang.org/pkg/os/#Open
	// io.Copy  https://golang.org/pkg/io/#Copy
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("I cannot open", fileName, "and got this error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
