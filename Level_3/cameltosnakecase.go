package piscine

func CamelToSnakeCase(s string) string{
	if len(s) == 0 {
		return ""
	}

	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}

	for i := 0; i < len(s); i++ {
		isUpper := s[i] >= 'A' && s[i] <= 'Z'
		isLower := s[i] >= 'a' && s[i] <= 'z'

		if !isUpper && ! isLower {
			return s
		}

		if i > 0 && (s[i-1] >= 'A' && s[i-1] <= 'Z') && isUpper {
			return s
		}
	}

	result := ""

	for i := 0; i < len(s); i++ {
		if i > 0 && (s[i] >= 'A' && s[i] <= 'Z') {
			result += "_"
		}
		result += string(s[i])
	}
	return result
}