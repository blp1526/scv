package main

import (
	"os"

	"github.com/blp1526/scv"
)

func main() {
	logger := &scv.Logger{}
	msg, err := scv.Run()
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	} else {
		logger.Info(msg)
	}
}
