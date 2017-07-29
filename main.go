package main

import (
	"fmt"
	"log"

	"github.com/blp1526/scv/cmd"
	"github.com/blp1526/scv/color"
)

func main() {
	log.SetFlags(0)
	msg, err := cmd.Run()
	if err != nil {
		log.Fatal(color.Red(fmt.Sprintf("fatal: %s", err)))
	} else {
		log.Println(msg)
	}
}
