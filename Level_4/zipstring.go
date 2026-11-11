package piscine

func ZipString(s string) string {
	i := 0
	result := ""
	for i < len(s) {
		count := 1
		for i+count < len(s) && s[i] == s[i+count] {
			count++
		}
		result += string(rune(count) + '0') + string(s[i])
		i += count
	}
	return result
}
