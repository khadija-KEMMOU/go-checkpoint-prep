package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	if len(slice2) > len(slice1) {
		slice1, slice2 = slice2, slice1
	}

	result := []int{}
	for i := 0; i < len(slice2); i++ {
		result = append(result, slice1[i])
		result = append(result, slice2[i])
	}


	result = append(result, slice1[len(slice2):]...)
	return result
}
