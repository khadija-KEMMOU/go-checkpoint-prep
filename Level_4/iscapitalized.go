package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		start := i == 0 || s[i-1] == ' '
		if start && s[i] >= 'a' && s[i] <= 'z' {
			return false
		}
	}

	return true
}
