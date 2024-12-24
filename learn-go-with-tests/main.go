package main

import (
	"os"

	mock "github.com/arunkhattri/ws-go/learn-go-with-tests/09-mocking"
)

func main() {
	num := 5
	dSleeper := &mock.DefaultSleeper{}
	mock.Countdown(os.Stdout, num, dSleeper)
}
