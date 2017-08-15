package scv

import (
	"fmt"
	"io"
)

type Logger struct {
	Verbose   bool
	OutStream io.Writer
	ErrStream io.Writer
}

func (logger *Logger) Debug(a string) (n int, err error) {
	if logger.Verbose {
		format, _ := logger.Format("lightGray", "debug")
		n, err = fmt.Fprintf(logger.OutStream, format, a)
		return n, err
	} else {
		return n, err
	}
}

func (logger *Logger) Info(a string) (n int, err error) {
	n, err = fmt.Fprintln(logger.OutStream, a)
	return n, err
}

// NOTE: (a interface{}) is expected to be string or error.
func (logger *Logger) Fatal(a interface{}) (n int, err error) {
	format, _ := logger.Format("red", "fatal")
	n, err = fmt.Fprintf(logger.ErrStream, format, a)
	return n, err
}

func (logger *Logger) Format(colorName string, level string) (format string, err error) {
	ansiColors := map[string]string{
		"normal":    "\033[m",
		"black":     "\033[30m",
		"red":       "\033[31m",
		"green":     "\033[32m",
		"yellow":    "\033[33m",
		"blue":      "\033[34m",
		"magenta":   "\033[35m",
		"cyan":      "\033[36m",
		"lightGray": "\033[37m",
	}
	ansiColor, hasKey := ansiColors[colorName]
	if !hasKey {
		err = fmt.Errorf("Unsupported colorName")
		return format, err
	}
	format = ansiColor + level + ": %s" + ansiColors["normal"] + "\n"
	return format, err
}
