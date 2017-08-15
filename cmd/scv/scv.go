package main

import (
	"os"

	"github.com/blp1526/scv"
)

func main() {
	cli := &scv.CLI{
		Logger: scv.Logger{
			OutStream: os.Stdout,
			ErrStream: os.Stderr,
		},
	}
	result, err := cli.Run()
	if err != nil {
		cli.Logger.Fatal(err)
		os.Exit(1)
	} else {
		cli.Logger.Info(result)
	}
}
