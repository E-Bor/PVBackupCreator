package utils

func SplitForNGroups[T string | int | float32 | float64](array []T, n int) [][]T {
	var splitedGroups [][]T
	length := len(array)
	offset := length / n
	remainder := length % n
	for i := 0; i+offset <= length; i += offset {
		splitedGroups = append(splitedGroups, array[i:i+offset])
	}
	if remainder > 0 {
		var subGroup []T
		for i := length - 1; i > length-remainder; i-- {
			subGroup = append(subGroup, array[i])
		}
		splitedGroups = append(splitedGroups, subGroup)
	}
	return splitedGroups
}
