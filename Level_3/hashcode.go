package piscine

func HashCode(dec string) string {
	result := ""
	for _, c := range dec {
		c = rune((int(c) + len(dec)) % 127)
		if c < 32 {
			c += 33
		}
		result += string(c)
	}
	return result
}
