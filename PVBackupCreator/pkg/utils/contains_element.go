package utils

func ContainsElement[T string | int | float32 | float64](array []T, element T) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}
