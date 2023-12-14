package main

import "github.com/01-edu/z01"

func main() {
	var i int = 0

	for i <= 9 {
		z01.PrintRune(rune('0' + i))
		i++
	}
	z01.PrintRune('\n')
}
