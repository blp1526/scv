package color

func Red(text string) string {
	return "\033[31m" + text + "\033[m"
}
