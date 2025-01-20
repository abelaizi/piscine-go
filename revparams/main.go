package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		args := os.Args[i]
		for _, char := range args {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
