package main

import "fmt"

func RenderEl(str string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", str)
}
