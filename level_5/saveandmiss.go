package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	save := ""
	for i := 0; i < len(arg); i++ {
		if (i/num)%2 == 0 {
			save += string(arg[i])
		}
	}

	return save
}
