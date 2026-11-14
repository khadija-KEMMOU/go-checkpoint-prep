package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	for _, arg := range os.Args[1:] {
		for i := 0; i < len(arg); i++ {
			c := arg[i]
			last := i == len(arg)-1 || arg[i+1] == ' '

			if last && c >= 'a' && c <= 'z' {
				c = c - 'a' + 'A'
			} else if !last && c >= 'A' && c <= 'Z' {
				c = c - 'A' + 'a'
			}
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
	}
}