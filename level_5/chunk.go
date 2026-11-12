package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

// with fmt func
func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}

	chunks := [][]int{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	fmt.Println(chunks)
}

// with z01 func
func Chunk2(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
	}

	z01.PrintRune('[')
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		if i > 0 {
			z01.PrintRune(' ')
		}
		z01.PrintRune('[')
		for j := i; j < end; j++ {
			if j > i {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(slice[j] + '0'))
		}
		z01.PrintRune(']')
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}
