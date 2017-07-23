package main

import (
	"fmt"
	"log"

	"github.com/blp1526/scv/cmd"
	"github.com/blp1526/scv/color"
)

func main() {
	msg, err := cmd.Run()
	if err != nil {
		log.SetFlags(0)
		log.Fatal(color.Red(fmt.Sprintf("fatal: %s", err)))
	} else {
		fmt.Println(msg)
	}
}
