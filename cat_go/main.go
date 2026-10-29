package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	// make a function which takes a path we return content of that file
	MyReadFile("test.txt")
}

func MyReadFile(path string) {
	if path == "" {
		fmt.Println(errors.New("need a file path"))

		return
	}

	// read the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		noOfBytesRead, err := file.Read(buffer)

		if noOfBytesRead > 0 {
			fmt.Println(string(buffer[:noOfBytesRead]))
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println(string(buffer[:noOfBytesRead]))
				return
			}
			fmt.Println(err)
			return
		}

	}
}
