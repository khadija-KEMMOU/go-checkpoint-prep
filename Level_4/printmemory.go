package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i++ {
		div := arr[i] / 16
		mod := arr[i] % 16
		hex := "0123456789abcdef"
		z01.PrintRune(rune(hex[div]))
		z01.PrintRune(rune(hex[mod]))
		z01.PrintRune(' ')
		if (i+1) % 4 == 0  {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
	for i := 0; i < len(arr); i++ {
		if arr[i] < 32 || arr[i] > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(arr[i]))
		}
	}
	z01.PrintRune('\n')
}
