package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	seen := ""
	count := 0
	for _, c := range str1 {
		if !strings.Contains(str2, string(c)) && !strings.Contains(seen, string(c)) {
			seen += string(c)
			count++
		}
	}

	for _, c := range str2 {
		if !strings.Contains(str1, string(c)) && !strings.Contains(seen, string(c)) {
			seen += string(c)
			count++
		}
	}

	return count
}
