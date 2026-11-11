package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, c := range s {
		x := 1
		if c >= 'a' && c <= 'z' {
			x = int(c - 'a') + 1
		} else if c >= 'A' && c <= 'Z' {
			x = int(c - 'A') + 1
		}
		for j := 0; j < x; j++ {
			result += string(c)
		}
	}
	return result
}
