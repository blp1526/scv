package color

func Red(msg string) string {
	return format("\033[31m", msg)
}

func format(color string, msg string) string {
	return color + msg + "\033[m"
}
