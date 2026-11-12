package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := []rune{}
	rev := ""

	if n < 0 {
		n = -n
		result = append(result, '-')
	}	

	for n > 0 {
		remainder := n % 10
		n = n / 10
		rev += string(rune(remainder + '0'))
	}

	for i := len(rev)-1; i >= 0; i-- {
		result = append(result, rune(rev[i]))
	}

	return string(result)
}