package logger

import (
	"fmt"
)

const normal = "\033[m"
const black = "\033[30m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const magenta = "\033[35m"
const cyan = "\033[36m"
const lightGray = "\033[37m"

func Debug(message string) {
	fmt.Printf("%sdebug%s: %s\n", lightGray, normal, message)
}

func Fatal(message string) {
	fmt.Printf("%sfatal%s: %s\n", red, normal, message)
}
