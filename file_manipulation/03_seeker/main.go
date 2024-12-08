package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// keyword
	keyword := "error"
	// open log file and check for errors if any
	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}

	// defer close
	defer f.Close()

	// file reader
	reader := bufio.NewReader(f)

	var lastKeywordPosition int64
	var currentPosition int64

	// read log file line by line
	for {
		line, err := reader.ReadString('\n')
		currentPosition += int64(len(line))
		// find the position of keyword if line contains keyword
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			// lastKeywordPosition = currentPosition - length of line
			lastKeywordPosition = currentPosition - int64(len(line))
			// fmt.Println("[L] ", line)
		}
		// handle all possible errors
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}
	}

	// seek from lastKeywordPosition
	if lastKeywordPosition != 0 {
		_, err = f.Seek(lastKeywordPosition, io.SeekStart)
		if err != nil {
			fmt.Println("error seeking file:", err)
			return
		}

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("error reading file:", err)
			return
		}
		fmt.Println(line)
	} else {
		fmt.Println("NOT FOUND")
	}
}
