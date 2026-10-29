package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// creating a basic wc unix utility in go so it helps user to info about file content like
	// word count,bytes,no of lines,no of words ,no of characters,file name with flag

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stdout, "%v", "Not enough arguments")
		os.Exit(1)
	}

	flag := os.Args[1]
	for _, arg := range os.Args[2:] {
		count, err := myWCUtility(flag, arg)
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v", "err")
			return
		}
		switch flag {
		case "-l":
			fmt.Printf("File %s has %d lines", arg, count)
		case "-w":
			fmt.Printf("File %s has %d words", arg, count)

		}
	}
}

func myWCUtility(flag, path string) (int, error) {
	// go run main.go -l <filepath>

	switch flag {
	case "-l":
		countOFLine, err := noOFLines(path)
		return countOFLine, err

	case "-w":
		wordCount, err := noOFWordINFile(path)
		return wordCount, err
	}
	return 0, nil
}

func noOFLines(path string) (int, error) {
	//  open the file
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buffer := make([]byte, 4096) // 4kb used mostly

	count := 0

	for {
		noOFBytesRead, err := file.Read(buffer)

		if noOFBytesRead > 0 {
			count += bytes.Count(buffer[:noOFBytesRead], []byte{'\n'})
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, err
		}
	}

	return count, nil
}

func noOFWordINFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	buffer := make([]byte, 4096)
	var lineBuffer []byte
	wdcount := 0
	for {
		noofBytesRead, err := file.Read(buffer)

		if noofBytesRead > 0 {
			for i := 0; i < noofBytesRead; i++ {
				ch := buffer[i]

				if ch == '\n' {
					wdcount += len(strings.Split(string(lineBuffer), " "))
					lineBuffer = lineBuffer[:0]
				} else {
					lineBuffer = append(lineBuffer, ch)
				}
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(lineBuffer) > 0 {
					wdcount += len(lineBuffer)
				}
				break
			}
			return 0, err
		}
	}
	return wdcount, nil
}
