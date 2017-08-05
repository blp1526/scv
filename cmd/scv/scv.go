package main

import (
	"os"

	"github.com/blp1526/scv"
)

func main() {
	cli := &scv.CLI{}
	result, err := cli.Run()
	if err != nil {
		cli.Logger.Fatal(err)
		os.Exit(1)
	} else {
		cli.Logger.Info(result)
	}
}
