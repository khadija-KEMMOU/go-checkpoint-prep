package piscine

func LastWord(s string) string{
	i := len(s)-1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	reverse := ""
	for i >= 0 && s[i] != ' ' {
		reverse += string(s[i])
		i--
	}

	result := ""
	for i := len(reverse)-1; i >= 0; i-- {
		result += string(reverse[i])
	}
	return result + "\n"
}



// func LastWord(s string) string {
// 	end := len(s) - 1
// 	for end >= 0 && s[end] == ' ' {
// 		end-- // skip trailing spaces
// 	}

// 	start := end
// 	for start >= 0 && s[start] != ' ' {
// 		start-- // move back to the space before the word
// 	}

// 	return s[start+1:end+1] + "\n"
// }
