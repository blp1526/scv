package scv

import "fmt"

const normal = "\033[m"
const black = "\033[30m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const magenta = "\033[35m"
const cyan = "\033[36m"
const lightGray = "\033[37m"

type Logger struct {
	Verbose bool
}

func (logger *Logger) Debug(a string) (n int, err error) {
	if logger.Verbose {
		n, err = fmt.Printf(format(lightGray, "debug"), a)
		return n, err
	} else {
		return n, err
	}
}

func (logger *Logger) Info(a string) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}

// NOTE: (a interface{}) is expected to be string or error.
func (logger *Logger) Fatal(a interface{}) (n int, err error) {
	n, err = fmt.Printf(format(red, "fatal"), a)
	return n, err
}

func format(color string, level string) string {
	return color + level + ": %s" + normal + "\n"
}
