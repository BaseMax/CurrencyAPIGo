package utils

func SliceContains[T comparable](s []T, i T) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}
