package main

import "github.com/01-edu/z01"

func quadE(x, y int) {
	for a := 1; a <= y; a++ {
		for b := 1; b <= x; b++ {
			if a == 1 && b == 1 {
				z01.PrintRune('A')
			} else if a == 1 && b > 1 && b < x {
				z01.PrintRune('B')
			} else if a == 1 && b == x {
				z01.PrintRune('C')
			} else if b == 1 && a > 1 && a < y {
				z01.PrintRune('B')
			} else if a > 1 && a < y && b > 1 && b < x {
				z01.PrintRune(' ')
			} else if a > 1 && a < y && b == x {
				z01.PrintRune('B')
			} else if a == y && b == 1 {
				z01.PrintRune('C')
			} else if a == y && b > 1 && b < x {
				z01.PrintRune('B')
			} else {
				z01.PrintRune('A')
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
	quadE(5, 3)
}
