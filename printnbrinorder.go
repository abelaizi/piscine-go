package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	runes := []rune{}
	if n == 0 {
		z01.PrintRune(0 + '0')
	}

	for n > 0 {
		runes = append(runes, '0'+rune(n%10))
		n /= 10
	}

	for i := '0'; i <= '9'; i++ {
		for _, r := range runes {
			if r == i {
				z01.PrintRune(r)
			}
		}
	}
}
