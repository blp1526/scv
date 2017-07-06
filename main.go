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
	argsSize := len(os.Args) - 1
	if argsSize == 0 {
		fmt.Println("scv version 0.0.1")
	} else if argsSize == 1 {
		err := cmd.Run("foo")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		msg := "ArgumentError: wrong number of arguments (given %d, expected 1)\n"
		fmt.Printf(msg, argsSize)
		os.Exit(1)
	}
}
