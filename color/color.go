package color

func Red(text string) string {
	return format("\033[31m", text)
}

func format(color string, text string) string {
	return color + text + "\033[m"
}
