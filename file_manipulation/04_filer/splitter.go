package main

import "os"

func SplitFile(filepath string, partSize int64) error {
	file, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// odd size
	totalParts := fileInfo.Size() / partSize
	if fileInfo.Size()%partSize != 0 {
		totalParts++
	}

	for {
	}

}
