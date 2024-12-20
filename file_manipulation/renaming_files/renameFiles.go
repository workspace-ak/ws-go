package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	// list all files to rename
	bhavcopies, err := filepath.Glob("./data/*.csv")
	if err != nil {
		log.Fatalln("Bad Pattern,", err)
	}

	// Check if filename contains "BhavCopy"
	for idx, bc := range bhavcopies {
		dir, file := filepath.Split(bc)
		matched, err := regexp.MatchString(`BhavCopy`, file)
		if err != nil {
			log.Fatalln("Incorrect regex;", err)
			return
		}
		if matched {
			re, err := regexp.Compile(`\d{8}`)
			if err != nil {
				log.Fatalln("regex not correct")
			}
			dt := re.FindString(file)
			newfname := filepath.Join(dir, dt+"_bhav.csv")
			os.Rename(bc, newfname)
		}
		if (idx+1)%5 == 0 {
			fmt.Printf("Renamed %d files...\n", idx+1)
		}
	}
}
