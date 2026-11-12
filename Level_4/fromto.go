package piscine

func FromTo(from int, to int) string {
	if (from < 0 || from > 99) || (to < 0 || to > 99) {
		return "Invalid\n"
	}

	step := 1
	if from > to {
		step = -1
	}
	result := ""
	for i := from; ; i += step {
		tens := i / 10
		ones := i % 10
		result += string(rune(tens + '0'))
		result += string(rune(ones + '0'))

		if i != to {
			result += ", "
		} else {
			break
		}
	}
	return result + "\n"
}
