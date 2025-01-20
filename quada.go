package main

import "github.com/01-edu/z01"

func quadA(x, y int) {
    for a := 1; a <= y; a++ {
        for b := 1; b <= x; b++ {
            if a == 1 && b == 1 {
                z01.PrintRune('o')
            } else if a == 1 && b > 1 && b < x {
                z01.PrintRune('-')
            } else if a == 1 && b == x {
                z01.PrintRune('o')
            } else if b == 1 && a > 1 && a < y {
                z01.PrintRune('|')
            } else if a > 1 && a < y && b > 1 && b < x {
                z01.PrintRune(' ')
            } else if a > 1 && a < y && b == x {
                z01.PrintRune('|')
            } else if a == y && b == 1 {
                z01.PrintRune('o')
            } else if a == y && b > 1 && b < x {
                z01.PrintRune('-')
            } else {
                z01.PrintRune('o')
            }
        }
        z01.PrintRune('\n')
    }
}
func main() {
    quadA(5, 3)
}
