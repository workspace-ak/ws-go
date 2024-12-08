package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data := "Welcome to a wonderful world of golang!"
	// alpha := []byte{65, 66, 67, 68, 69, 70}
	// fmt.Printf("%c\n", alpha)
	f, err := os.Create("./alpha.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// _, err = f.Write(alpha)
	_, err = f.Write([]byte(data))

	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Open("./alpha.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := io.ReadAll(f2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(bs))
}
