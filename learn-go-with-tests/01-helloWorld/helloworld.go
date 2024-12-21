package main

import (
	"fmt"
	"reflect"
)

type greet string

const (
	engGreetPrefix    greet = "Hello"
	hindiGreetPrefix  greet = "Namaskar"
	frenchGreetPrefix greet = "Bonjour"
)

// empty values will be replaced with defaults
type Parameters struct {
	name string `default:"World"`
	l    string `default:"english"`
}

func Hello(prm Parameters) string {
	typ := reflect.TypeOf(prm)
	if prm.name == "" {
		f, _ := typ.FieldByName("name")
		prm.name = f.Tag.Get("default")
	}
	if prm.l == "" {
		f, _ := typ.FieldByName("l")
		prm.l = f.Tag.Get("default")
	}

	var prefix greet

	switch prm.l {
	case "hindi":
		prefix = hindiGreetPrefix
	case "french":
		prefix = frenchGreetPrefix
	default:
		prefix = engGreetPrefix
	}
	return fmt.Sprintf("%s, %s", prefix, prm.name)
}

func main() {
	param := Parameters{"", ""}
	fmt.Println(Hello(param))
}
