package goutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Read line by line from a TXT file,
// then returns it as string list.
func File2List(filepath string) (items []string, err error) {
	// filepath := os.Args[1]
	readFile, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		items = append(items, fileScanner.Text())
	}

	readFile.Close()

	return
}

func List2File(filepath string, values []interface{}) {
	file, err := os.Create(filepath)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Defer is used for purposes of cleanup like
	// closing a running file after the file has
	// been written and main //function has
	// completed execution
	defer file.Close()

	// len variable captures the length
	// of the string written to the file.
	len, err := file.WriteString(strings.Join(ToStringArray(values), "\n"))

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// Name() method returns the name of the
	// file as presented to Create() method.
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
