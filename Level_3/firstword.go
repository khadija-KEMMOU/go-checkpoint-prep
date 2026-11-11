package piscine

func FirstWord(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	result := ""
    for i < len(s) && s[i] != ' ' {
		result += string(s[i])
		i++
	}
	
	return result + "\n"
}
