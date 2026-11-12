package piscine

func FindPrevPrime(nb int) int {
	for nb >= 2 {
		isprime := true
		for i := 2; i < nb; i++ {
			if nb % i == 0 {
				isprime = false
				break
			}
		}

		if isprime {
			return nb
		} else {
			nb--
		}
	}
	return 0
}
