package main

import (
	"fmt"

	"github.com/thiagozs/go-logruswr"
)

func main() {
	log, _ := logruswr.New()

	log.Debug("Debug test")
	log.Info("Info test")
	log.Warn("Warning test")
	log.Error("Error test")

	log.WithError(fmt.Errorf("teste")).Info("Error with error test")

	log.WithFields(logruswr.Fields{
		"test":  "test",
		"test2": "test2",
	}).Info("WithFields test")

	log.Fatal("Fatal test")
}
