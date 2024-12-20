package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// get all the relevant files
	bhavcopies, err := filepath.Glob("./data/*.csv")
	if err != nil {
		log.Fatalln("Bad Pattern,", err)
	}

	// merge files
	mergedfile, err := os.Create("./data/merged_bhav.csv")
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer mergedfile.Close()

	writer := csv.NewWriter(mergedfile)
	defer writer.Flush()

	for i, bc := range bhavcopies {
		f, err := os.Open(bc)
		if err != nil {
			log.Fatalln("Error:", err)
			return
		}
		defer f.Close()
		reader := csv.NewReader(f)
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatalln("Error reading file:", err)
			return
		}
		if i > 0 {
			records = records[1:]
		}
		for _, record := range records {
			if err := writer.Write(record); err != nil {
				log.Fatalln("Error writing record to file:", err)
			}
		}
		fmt.Printf("Merged %d files\n", i+1)
	}
}
