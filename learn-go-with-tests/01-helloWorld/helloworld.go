package main

import (
	"fmt"
)

// separate 'domain' code

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(Hello("world"))
}
