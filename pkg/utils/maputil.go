package utils

func MapKeyToSlice[T comparable](m map[T]T) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
