package helloworld

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
	Name string `default:"World"`
	Lang string `default:"english"`
}

func Hello(prm Parameters) string {
	typ := reflect.TypeOf(prm)
	if prm.Name == "" {
		f, _ := typ.FieldByName("Name")
		prm.Name = f.Tag.Get("default")
	}
	if prm.Lang == "" {
		f, _ := typ.FieldByName("Lang")
		prm.Lang = f.Tag.Get("default")
	}

	var prefix greet

	switch prm.Lang {
	case "hindi":
		prefix = hindiGreetPrefix
	case "french":
		prefix = frenchGreetPrefix
	default:
		prefix = engGreetPrefix
	}
	return fmt.Sprintf("%s, %s", prefix, prm.Name)
}
