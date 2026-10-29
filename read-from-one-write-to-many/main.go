package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// user should give me <srcfile path> <destfile path1> <destfile path2>
	a := time.Now()
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stdout, "%s", "no args provided")
	}
	var wg sync.WaitGroup
	for _, arg := range os.Args[2:] {
		wg.Go(func() {
			MyFileWriter(os.Args[1], arg)
		})
	}
	wg.Wait()

	fmt.Println(time.Since(a))
}

func MyFileWriter(srcPath string, destPath string) error {
	// open the file
	srFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	defer srFile.Close()

	if len(strings.Split(destPath, ".")) != 2 {
		return errors.New("bad input try again bruh")
	}
	dsFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dsFile.Close()

	buffer := make([]byte, 4096)
free:
	for {

		NoOfBytesRead, err := srFile.Read(buffer)

		if NoOfBytesRead > 0 {
			if _, err := dsFile.Write(buffer[:NoOfBytesRead]); err != nil {
				return err
			}
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
