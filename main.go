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
	fmt.Printf("test %s\n", os.Args[1])
}
