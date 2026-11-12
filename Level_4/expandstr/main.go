package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	text := os.Args[1]
	started := false
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' && text[i] != '\t' {
			if started && (text[i-1] == ' ' || text[i-1] == '\t') {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(text[i]))
			started = true
		}
	}

	if started {
		z01.PrintRune('\n')
	}
}