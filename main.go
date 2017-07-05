package main

import (
	"fmt"
	"os"
)

func main() {
	argsSize := len(os.Args) - 1
	if argsSize != 1 {
		msg := "ArgumentError: wrong number of arguments (given %d, expected 1)\n"
		fmt.Printf(msg, argsSize)
		os.Exit(1)
	}

	// TODO: This app's binary is MacOS only.
	//
	// * Open $HOME/scv.yml
	//   * then get Web API ACCESS_TOKEN and ACCESS_TOKEN_SECRET
	//   * then convert os.Args[1] to Web API :id
	// * Request "GET /somewhere/:id", then get response as JSON format
	// * Make "open command argument" from JSON
	// * Run MacOS command "open argument"
	fmt.Printf("open %s\n", os.Args[1])
}
