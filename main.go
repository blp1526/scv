package main

import (
	"os"

	"github.com/blp1526/scv/cmd"
	"github.com/blp1526/scv/logger"
)

func main() {
	l := &logger.Logger{}
	msg, err := cmd.Run()
	if err != nil {
		l.Fatal(err)
		os.Exit(1)
	} else {
		l.Info(msg)
	}
}
