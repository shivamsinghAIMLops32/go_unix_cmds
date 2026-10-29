package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	// lets take argument we expect 1 argument
	if len(os.Args) < 1 {
		// cause we expect go run main.go <path> so
		fmt.Fprintf(os.Stdout, "%v", "we need a argument")
		os.Exit(1)
	}

	// the first arg os.args return is the path of running program so like
	// go run main.go 12 ab AB  => []string{main.go,12,ab,AB}
	for _, arg := range os.Args[1:] {
		if err := MyFileReader(arg); err != nil {
			fmt.Fprintf(os.Stdout, "%s", err)
		}
	}
}

func MyFileReader(path string) error {
	// open the file at the path
	file, err := os.Open(path)
	if err != nil {
		return errors.New("error in opening the file")
	}

	defer file.Close()

	// keep reading the content of file unless you are at the io.EOF
	buffer := make([]byte, 4096) // 4kb is standard used
free:
	for {
		// number of bytes read tells how many new bytes file.Read has wrote in buffer
		noOfBytesRead, err := file.Read(buffer)

		if noOfBytesRead > 0 {
			os.Stdout.Write(buffer)
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break free
			}
			return err
		}
	}
	return nil
}
