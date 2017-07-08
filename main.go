// This app's binary is MacOS only.
//
// * Open $HOME/scv.json
//   * then get Web API ACCESS_TOKEN and ACCESS_TOKEN_SECRET
//   * then convert os.Args[1] to Web API :id
// * Request "GET /somewhere/:id", then get response as JSON format
// * Make "open command argument" from JSON
// * Run MacOS command "open argument"
package main

import (
	"fmt"
	"os"

	"github.com/blp1526/scv/cmd"
)

func main() {
	version := "0.0.1"
	expectedArgsSize := 2

	argsSize := len(os.Args) - 1
	if argsSize == 0 {
		fmt.Printf("scv version %s\n", version)
	} else if argsSize == 2 {
		err := cmd.Run(os.Args[1], os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		msg := "fatal: Wrong number of arguments (given %d, expected %d)\n"
		fmt.Printf(msg, argsSize, expectedArgsSize)
		os.Exit(1)
	}
}
