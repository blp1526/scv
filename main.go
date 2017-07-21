package main

import (
	"fmt"
	"os"

	"github.com/blp1526/scv/cmd"
	"github.com/blp1526/scv/logger"
)

func main() {
	msg, err := cmd.Run()
	if err != nil {
		logger.Fatal(fmt.Sprintf("%s", err))
		os.Exit(1)
	} else {
		fmt.Println(msg)
	}
}
