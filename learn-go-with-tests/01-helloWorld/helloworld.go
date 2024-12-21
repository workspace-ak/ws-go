package main

import (
	"fmt"
)

type greet string

const (
	engGreetPrefix    greet = "Hello"
	hindiGreetPrefix  greet = "Namaskar"
	frenchGreetPrefix greet = "Bonjour"
)

type lang string

const (
	english lang = "English"
	hindi   lang = "Hindi"
	french  lang = "French"
)

// separate 'domain' code

func Hello(name string, l lang) string {

	if name == "" {
		name = "World"
	}
	var prefix greet

	switch l {
	case hindi:
		prefix = hindiGreetPrefix
	case french:
		prefix = frenchGreetPrefix
	default:
		prefix = engGreetPrefix
	}
	return fmt.Sprintf("%s, %s", prefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
