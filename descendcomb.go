package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for b := '9'; b >= '1'; b-- {
			d := b - 1
			for c := i; c >= '0'; c-- {
				for ; d >= '0'; d-- {

					z01.PrintRune(i)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)

					if i > '0' || b > '1' || c > '0' || d > '0' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				d = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
