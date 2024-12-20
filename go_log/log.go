package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("This is a log message")
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("User logged in",
		"username", "arun",
		"userid", 123,
		"login_attempt", 1,
	)
}
