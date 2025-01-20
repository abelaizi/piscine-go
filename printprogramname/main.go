package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := os.Args[0][2:]
	for _, z := range i {
		z01.PrintRune(z)
	}
	z01.PrintRune('\n')
}
