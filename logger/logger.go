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

func Debug(msg string) {
	fmt.Println(format(lightGray, "debug", msg))
}

func Fatal(msg string) {
	fmt.Println(format(red, "fatal", msg))
}

func format(color string, level string, msg string) string {
	return fmt.Sprintf("%s%s%s: %s", color, level, normal, msg)
}
