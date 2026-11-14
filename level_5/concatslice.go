package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	result := []int{}
	result = append(result, slice1...)
	result = append(result, slice2...)
	return result
}
