package main

import (
	"fmt"
	"os"

	"github.com/blp1526/scv/cmd"
	"github.com/blp1526/scv/logger"
)

const version = "0.0.1"
const expectedArgsSize = 2

func main() {
	argsSize := len(os.Args) - 1
	if argsSize == 0 {
		fmt.Printf("scv version %s\n", version)
	} else if argsSize == expectedArgsSize {
		zoneName := os.Args[1]
		serverName := os.Args[2]
		err := cmd.Run(zoneName, serverName)
		if err != nil {
			logger.Fatal(fmt.Sprintf("%s", err))
		}
	} else {
		logger.Fatal(fmt.Sprintf("Expected arguments size is %d", expectedArgsSize))
	}
}
