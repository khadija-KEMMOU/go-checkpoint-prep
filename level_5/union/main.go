package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
	}
	
	text := os.Args[1] + os.Args[2]
	seen := ""

	for _, c := range text {
		found := false
		for _, s := range seen {
			if c == s {
				found = true
			}
		}

		if !found {
			seen += string(c)
		}
	}

	for _, s := range seen {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}